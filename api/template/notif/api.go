package notif

import (
	"github.com/NpoolPlatform/message/npool/third/mgr/v1/template/notif"
	"google.golang.org/grpc"
)

type Server struct {
	notif.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	notif.RegisterManagerServer(server, &Server{})
}
