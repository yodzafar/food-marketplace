package app

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/yodzafar/food-market/api-gateway/config"
	"github.com/yodzafar/food-market/api-gateway/internal/middleware"
	pb "github.com/yodzafar/food-marketpalce/gen/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	cfg    *config.Config
	server *http.Server
}

func New(cfg *config.Config) (*App, error) {
	ctx := context.Background()
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	if err := pb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, cfg.Services.UserAddr, opts); err != nil {
		return nil, err
	}

	handler := middleware.Logger(middleware.CORS(mux))

	server := &http.Server{
		Addr:    ":" + cfg.HTTP.Port,
		Handler: handler,
	}

	return &App{cfg: cfg, server: server}, nil
}

func (a *App) Run() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("api-gateway started on port %s", a.cfg.HTTP.Port)
		if err := a.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	<-quit
	log.Println("api-gateway shutting down")
	if err := a.server.Shutdown(context.Background()); err != nil {
		log.Fatalf("server shutdown error: %v", err)
	}
	log.Println("api-gateway stopped")
}
