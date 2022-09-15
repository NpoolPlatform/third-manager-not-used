package api

import (
	"github.com/NpoolPlatform/message/npool/servicetmpl"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"github.com/NpoolPlatform/third-manager/api/contact"
)

type Server struct {
	servicetmpl.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	servicetmpl.RegisterManagerServer(server, &Server{})
	contact.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
