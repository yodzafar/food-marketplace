package app

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/yodzafar/food-marketpalce/user-service/config"
	grpcdelivery "github.com/yodzafar/food-marketpalce/user-service/internal/delivery/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	grpcServer *grpc.Server
	cfg        *config.Config
}

func New(cfg *config.Config, handler *grpcdelivery.UserHandler) *App {
	s := grpc.NewServer()
	handler.RegisterService(s)
	reflection.Register(s)
	return &App{
		grpcServer: s,
		cfg:        cfg,
	}
}

func (a *App) Run() {
	lis, err := net.Listen("tcp", ":"+a.cfg.GRPC.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("user-service started on port %s", a.cfg.GRPC.Port)
		if err := a.grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	<-quit
	log.Println("user-service shutting down")
	a.grpcServer.GracefulStop()
	log.Println("user-service stopped")
}
