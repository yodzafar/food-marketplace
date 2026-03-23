package grpcdelivery

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewUserHandler)
