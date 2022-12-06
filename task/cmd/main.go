package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
	"task/config"
	"task/discovery"
	"task/internal/dao"
	"task/internal/handler"
	"task/internal/service"
)

func main() {
	config.InitConfig()
	dao.InitDB()

	//取出etcd的地址
	etcdAddress := []string{viper.GetString("etcd.address")}
	//服务的注册
	etcdRegsiter := discovery.NewRegister(etcdAddress, logrus.New())
	grpcAddress := viper.GetString("server.grpcAddress")
	userNode := discovery.Server{
		Name:    viper.GetString("server.domain"),
		Address: grpcAddress,
	}
	server := grpc.NewServer()
	defer server.Stop()
	//绑定
	service.RegisterTaskServiceServer(server, handler.NewTaskService())
	lis, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		panic(err)
	}
	if _, err := etcdRegsiter.Register(userNode, 10); err != nil {
		panic(fmt.Sprintf("start server failed, err: %v", err))
	}
	logrus.Info("server started listen on ", grpcAddress)
	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}
