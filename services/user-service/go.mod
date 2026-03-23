module github.com/yodzafar/food-marketpalce/user-service

go 1.26.1

require (
	github.com/Masterminds/squirrel v1.5.4
	github.com/google/uuid v1.6.0
	github.com/google/wire v0.7.0
	github.com/jmoiron/sqlx v1.4.0
	github.com/joho/godotenv v1.5.1
	github.com/lib/pq v1.11.2
	github.com/redis/go-redis/v9 v9.18.0
	github.com/yodzafar/food-marketpalce/gen/proto v0.0.0
	golang.org/x/crypto v0.49.0
	google.golang.org/grpc v1.79.2
	google.golang.org/protobuf v1.36.11
)

replace github.com/yodzafar/food-marketpalce/gen/proto => ../../gen/proto

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.28.0 // indirect
	github.com/lann/builder v0.0.0-20180802200727-47ae307949d0 // indirect
	github.com/lann/ps v0.0.0-20150810152359-62de8c46ede0 // indirect
	github.com/stretchr/testify v1.5.1 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	golang.org/x/net v0.52.0 // indirect
	golang.org/x/sys v0.42.0 // indirect
	golang.org/x/text v0.35.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20260311181403-84a4fc48630c // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260311181403-84a4fc48630c // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)
