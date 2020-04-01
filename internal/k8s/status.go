package k8s

import (
	"context"
	"fmt"
	"strings"

	"github.com/golang/protobuf/ptypes/timestamp"
	modoki "github.com/modoki-paas/modoki-k8s/api"
	"golang.org/x/xerrors"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func k8sTimeToGRPCTimestamp(t metav1.Time) *timestamp.Timestamp {
	return &timestamp.Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.Nanosecond()),
	}
}

func convertContainerStatus(podStat *v1.PodStatus, containerName string) *modoki.ContainerStatus {
	var contStat *v1.ContainerStatus

	for i := range podStat.ContainerStatuses {
		if podStat.ContainerStatuses[i].Name == containerName {
			contStat = &podStat.ContainerStatuses[i]

			break
		}
	}

	ret := &modoki.ContainerStatus{
		Phase:   modoki.ContainerPhase(modoki.ContainerPhase_value[string(podStat.Phase)]),
		Message: podStat.Message,
		Reason:  podStat.Reason,
	}

	if contStat == nil {
		return ret
	}

	ret.ImageId = contStat.ImageID
	ret.RestartCount = contStat.RestartCount
	ret.Ready = contStat.Ready

	switch {
	case contStat.State.Running != nil:
		ret.StartedAt = k8sTimeToGRPCTimestamp(contStat.State.Running.StartedAt)
	case contStat.State.Terminated != nil:
		term := contStat.State.Terminated

		ret.ExitCode = term.ExitCode
		ret.Signal = term.Signal

		ret.StartedAt = k8sTimeToGRPCTimestamp(term.StartedAt)
		ret.FinishedAt = k8sTimeToGRPCTimestamp(term.FinishedAt)

		ret.Message = ret.Message + " " + term.Message
		ret.Reason = ret.Message + " " + term.Reason
	case contStat.State.Waiting != nil:
		waiting := contStat.State.Waiting

		ret.Message = ret.Message + " " + waiting.Message
		ret.Reason = ret.Message + " " + waiting.Reason
	}

	ret.Reason = strings.TrimSpace(ret.Reason)
	ret.Message = strings.TrimSpace(ret.Message)

	return ret
}

// Status returns status for deployment and pods created by it
func (c *Client) Status(ctx context.Context, namespace, deploy, containerName string) (*modoki.AppStatus, error) {
	stat := &modoki.AppStatus{}

	dpl, err := c.clientset.AppsV1().
		Deployments(namespace).
		Get(deploy, metav1.GetOptions{})

	if err != nil {
		return nil, xerrors.Errorf("failed to get deployment for %s/%s: %w", namespace, deploy, err)
	}

	dplStat := dpl.Status

	stat.Available = dplStat.AvailableReplicas
	stat.Ready = dplStat.ReadyReplicas
	stat.Existing = dplStat.Replicas

	selector := metav1.FormatLabelSelector(dpl.Spec.Selector)

	pods, err := c.clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{
		LabelSelector: selector,
	})

	if err != nil {
		return nil, xerrors.Errorf("failed to list pods in deployment(%s) in namespace(%s): %w", deploy, namespace, err)
	}

	stat.Containers = make([]*modoki.ContainerStatus, len(pods.Items))
	for i := range pods.Items {
		stat.Containers[i] = convertContainerStatus(&pods.Items[i].Status, containerName)
	}

	var maxReplicas int32
	if dpl.Spec.Replicas != nil {
		maxReplicas = *dpl.Spec.Replicas
	} else {
		maxReplicas = 1
	}

	if maxReplicas == stat.Ready {
		stat.State = fmt.Sprintf("Running(%d/%d)", stat.Ready, maxReplicas)
	} else {
		hasError := false
		for i := range stat.Containers {
			if !stat.Containers[i].Ready && stat.Containers[i].Reason != "" {
				hasError = true
				break
			}
		}

		if hasError {
			stat.State = fmt.Sprintf("Error(%d/%d)", stat.Ready, maxReplicas)
		} else {
			stat.State = fmt.Sprintf("Updating(%d/%d)", stat.Ready, maxReplicas)
		}
	}

	return stat, nil
}
