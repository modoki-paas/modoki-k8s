package auth

import (
	"context"
	"strings"

	"google.golang.org/grpc/metadata"
)

// Handle external requests

// GetToken retrieves token from HTTP headers
func GetToken(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return ""
	}

	headers := md.Get("Authorization")

	key := strings.TrimPrefix(headers[0], "Bearer ")

	return key
}

// GetTargetID returns the target to exec API for
func GetTargetID(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return ""
	}

	arr := md.Get(TargetIDHeader)

	if len(arr) == 0 {
		return ""
	}

	if !ok {
		return ""
	}

	return arr[0]
}
