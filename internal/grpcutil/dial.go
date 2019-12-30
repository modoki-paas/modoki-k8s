package grpcutil

import "google.golang.org/grpc"

func Dial(endpoint string, insecure bool) (*grpc.ClientConn, error) {
	opts := []grpc.DialOption{}

	if insecure {
		opts = append(opts, grpc.WithInsecure())
	}

	return grpc.Dial(endpoint, opts...)
}
