package api

import (
	"context"

	v1 "github.com/NpoolPlatform/message/npool/third/mgr/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	v1.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	v1.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return v1.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
