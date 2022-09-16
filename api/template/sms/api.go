package sms

import (
	"github.com/NpoolPlatform/message/npool/third/mgr/v1/template/sms"
	"google.golang.org/grpc"
)

type Server struct {
	sms.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	sms.RegisterManagerServer(server, &Server{})
}
