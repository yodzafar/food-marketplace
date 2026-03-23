module github.com/yodzafar/food-market/api-gateway

go 1.26.1

require (
	github.com/google/wire v0.7.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.28.0
	github.com/yodzafar/food-marketpalce/gen/proto v0.0.0
	google.golang.org/grpc v1.79.2
)

replace github.com/yodzafar/food-marketpalce/gen/proto => ../gen/proto

require (
	golang.org/x/net v0.52.0 // indirect
	golang.org/x/sys v0.42.0 // indirect
	golang.org/x/text v0.35.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20260311181403-84a4fc48630c // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260311181403-84a4fc48630c // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)
