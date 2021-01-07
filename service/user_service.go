package service

import (
	"github.com/VulpesFerrilata/grpc/protoc/user"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/client/grpc"
)

func NewUserService(opts ...client.Option) user.UserService {
	service := micro.NewService(
		micro.Name("boardgame.user.svc.client"),
		micro.Version("latest"),
		micro.Client(
			grpc.NewClient(
				opts...,
			),
		),
	)

	return user.NewUserService("boardgame.user.svc", service.Client())
}
