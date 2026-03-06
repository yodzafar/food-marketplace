package wire

import (
	"github.com/google/wire"
	"github.com/yodzafar/food-marketpalce/user-service/config"
	"github.com/yodzafar/food-marketpalce/user-service/internal/pkg/postgres"
	"github.com/yodzafar/food-marketpalce/user-service/internal/pkg/redis"
)

func initApp() (error, error) {

	wire.Build(
		config.LoadConfig(),
		postgres.ProviderSet,
		redis.ProviderSet,
	)

	return nil, nil
}
