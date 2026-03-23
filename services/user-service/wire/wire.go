//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/yodzafar/food-marketpalce/user-service/config"
	"github.com/yodzafar/food-marketpalce/user-service/internal/app"
	grpcdelivery "github.com/yodzafar/food-marketpalce/user-service/internal/delivery/grpc"

	"github.com/yodzafar/food-marketpalce/user-service/internal/pkg/postgres"
	redispkg "github.com/yodzafar/food-marketpalce/user-service/internal/pkg/redis"
	"github.com/yodzafar/food-marketpalce/user-service/internal/repository"
	"github.com/yodzafar/food-marketpalce/user-service/internal/usecase"
)

func InitApp() (*app.App, error) {
	wire.Build(
		config.LoadConfig,
		postgres.ProviderSet,
		redispkg.ProviderSet,
		repository.ProviderSet,
		usecase.ProviderSet,
		grpcdelivery.ProviderSet,
		app.New,
	)
	return nil, nil
}
