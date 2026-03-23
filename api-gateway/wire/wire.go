//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/yodzafar/food-market/api-gateway/config"
	"github.com/yodzafar/food-market/api-gateway/internal/app"
)

func initApp() (*app.App, error) {
	wire.Build(
		config.Load,
		app.New,
	)

	return nil, nil
}
