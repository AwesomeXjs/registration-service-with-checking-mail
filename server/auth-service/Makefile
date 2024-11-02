LOCAL_BIN:=$(CURDIR)/bin

run:
	go run cmd/auth/main.go

debug:
	go run cmd/auth/main.go -l debug

install-deps:
	mkdir -p $(LOCAL_BIN)
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1 # protofiles
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 # protofiles
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0 # migrations
	GOBIN=$(LOCAL_BIN) go install github.com/gojuno/minimock/v3/cmd/minimock@v3.4.0 # mocks
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.15.2 #gateway
	GOBIN=$(LOCAL_BIN) go install github.com/bojand/ghz/cmd/ghz@latest # для нагрузки по grpc
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.15.2 # swagger


get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go # protogen
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc # protogen
	go get google.golang.org/grpc # grpc
	go get google.golang.org/grpc/metadata # grpc metadata
	go get google.golang.org/grpc/internal/transport@v1.67.1 # grpc
	go get github.com/grpc-ecosystem/go-grpc-middleware # grpc middlewares
	go get -u github.com/jackc/pgx/v4 # pg driver
	go get github.com/jackc/pgx/v4/pgxpool@v4.18.3 # pg pool
	go get github.com/georgysavva/scany/pgxscan # pg actions
	go get -u github.com/brianvoe/gofakeit # fake data
	go get -u github.com/Masterminds/squirrel # query builder
	go get github.com/joho/godotenv # env config
	go get github.com/pkg/errors # errors
	go get github.com/stretchr/testify/require # tests
	go get github.com/golang/mock/gomock # mocks
	go get github.com/rs/cors # cors
	go get golang.org/x/crypto/bcrypt # hash password
	go get github.com/golang-jwt/jwt # jwt
	go get go.uber.org/zap # logger
	go get go.uber.org/zap/zapcore # logger
	go get github.com/natefinch/lumberjack # logger (for Zap)
	go get github.com/prometheus/client_golang/prometheus # prometheus
	go get github.com/prometheus/client_golang/prometheus/promauto # prometheus
	go get github.com/prometheus/client_golang/prometheus/promhttp@v1.20.4 # prometheus
	go get github.com/opentracing/opentracing-go # opentracing
	go get github.com/opentracing/opentracing-go/ext # opentracing
	go get github.com/uber/jaeger-client-go # tracing
	go get github.com/uber/jaeger-client-go/config # tracing config
	go get github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc # tracing (для мидлварь при создании клиента)
	go get github.com/sony/gobreaker # circuit breaker pattern
	go get github.com/fatih/color # terminal colors
	go get github.com/redis/go-redis/v9 # redis
	go get github.com/goccy/go-json # json decoder


# файлы для конфортной работы с плагинами grpc
vendor-proto:
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi
		@if [ ! -d vendor.protogen/protoc-gen-openapiv2 ]; then \
			mkdir -p vendor.protogen/protoc-gen-openapiv2/options &&\
			git clone https://github.com/grpc-ecosystem/grpc-gateway vendor.protogen/openapiv2 &&\
			mv vendor.protogen/openapiv2/protoc-gen-openapiv2/options/*.proto vendor.protogen/protoc-gen-openapiv2/options &&\
			rm -rf vendor.protogen/openapiv2 ;\
		fi

generate:
	make generate-auth-api

generate-auth-api:
	mkdir -p pkg/auth_v1
	protoc --proto_path api/auth_v1 --proto_path vendor.protogen \
    --go_out=pkg/auth_v1 --go_opt=paths=source_relative \
    --plugin=protoc-get-go=bin/protoc-gen-go \
    --go-grpc_out=pkg/auth_v1 --go-grpc_opt=paths=source_relative \
    --plugin=protoc-get-go-grpc=bin/protoc-gen-go-grpc \
	--grpc-gateway_out=pkg/auth_v1 --grpc-gateway_opt=paths=source_relative \
	--plugin=protoc-gen-grpc-gateway=bin/protoc-gen-grpc-gateway \
	--openapiv2_out=allow_merge=true,merge_file_name=api:pkg/swagger \
    --plugin=protoc-gen-openapiv2=bin/protoc-gen-openapiv2 \
    api/auth_v1/auth.proto