package api

import (
	"context"

	"github.com/NpoolPlatform/third-manager/api/template/frontend"

	"github.com/NpoolPlatform/third-manager/api/template/email"
	"github.com/NpoolPlatform/third-manager/api/template/sms"

	v1 "github.com/NpoolPlatform/message/npool/third/mgr/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"github.com/NpoolPlatform/third-manager/api/contact"
)

type Server struct {
	v1.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	v1.RegisterManagerServer(server, &Server{})
	contact.Register(server)
	email.Register(server)
	sms.Register(server)
	frontend.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return v1.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
