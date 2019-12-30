package k8s

import (
	"os"

	"golang.org/x/xerrors"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type Client struct {
	kubeconfig  string
	clientset   *kubernetes.Clientset
	kubectlPath string
}

// NewClient retunrs kubernetes client for client-go and kubectl
// kubeconfig must be path to file or empty(in-cluster config will be used)
func NewClient(kubeconfig string) (*Client, error) {
	var k8sConfig *rest.Config
	var err error

	if kubeconfig == "" {
		k8sConfig, err = rest.InClusterConfig()

		if err != nil {
			return nil, xerrors.Errorf("failed to initialize in cluster config: %w", err)
		}
	} else {
		k8sConfig, err = clientcmd.BuildConfigFromFlags("", kubeconfig)

		if err != nil {
			return nil, xerrors.Errorf("failed to initialize config from kubernetes: %w", err)
		}

	}

	clientset, err := kubernetes.NewForConfig(k8sConfig)

	if err != nil {
		return nil, xerrors.Errorf("failed to initialize client-go client: %w", err)
	}

	kubectlPath := "kubectl"

	if kp, ok := os.LookupEnv("KUBECTL_PATH"); ok {
		kubectlPath = kp
	}

	return &Client{
		clientset:   clientset,
		kubeconfig:  kubeconfig,
		kubectlPath: kubectlPath,
	}, nil
}
