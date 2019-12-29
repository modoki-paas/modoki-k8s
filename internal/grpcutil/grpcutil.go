package grpcutil

import (
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
)

func GRPCTimestamp(t time.Time) *timestamp.Timestamp {
	return &timestamp.Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.UnixNano() - t.Unix()*1e9),
	}
}
