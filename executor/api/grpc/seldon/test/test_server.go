package test

import (
	"context"
	"github.com/go-logr/logr"
	"github.com/seldonio/seldon-core/executor/api/grpc/seldon/proto"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"time"
)

type GrpcSeldonTestServer struct {
	log   logr.Logger
	delay int
}

func NewSeldonTestServer(delay int) *GrpcSeldonTestServer {
	return &GrpcSeldonTestServer{
		log:   logf.Log.WithName("GrpcSeldonTestServer"),
		delay: delay,
	}
}

func (g GrpcSeldonTestServer) Predict(ctx context.Context, req *proto.SeldonMessage) (*proto.SeldonMessage, error) {
	if g.delay > 0 {
		time.Sleep(time.Duration(g.delay) * time.Second)
	}
	return req, nil
}

func (g GrpcSeldonTestServer) SendFeedback(ctx context.Context, req *proto.Feedback) (*proto.SeldonMessage, error) {
	panic("Not implemented")
}
