package contact

import (
	"github.com/NpoolPlatform/message/npool/third/mgr/v1/contact"
	_ "github.com/NpoolPlatform/message/npool/third/mgr/v1/contact"
	"google.golang.org/grpc"
)

type Server struct {
	contact.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	contact.RegisterManagerServer(server, &Server{})
}
