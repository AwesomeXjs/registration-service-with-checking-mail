LOCAL_BIN:=$(CURDIR)/bin

run:
	go run cmd/mail/main.go

debug:
	go run cmd/mail/main.go -l debug

lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml

install-deps:
	mkdir -p $(LOCAL_BIN)
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1 # protofiles
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 # protofiles
	GOBIN=$(LOCAL_BIN) go install github.com/gojuno/minimock/v3/cmd/minimock@v3.4.0 # mocks
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.60.3 # lint


get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go # protogen
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc # protogen
	go get google.golang.org/grpc # grpc
	go get google.golang.org/grpc/metadata # grpc metadata
	go get google.golang.org/grpc/internal/transport@v1.67.1 # grpc
	go get github.com/grpc-ecosystem/go-grpc-middleware # grpc middlewares
	go get -u github.com/brianvoe/gofakeit # fake data
	go get github.com/joho/godotenv # env config
	go get github.com/pkg/errors # errors
	go get github.com/stretchr/testify/require # tests
	go get github.com/golang/mock/gomock # mocks
	go get go.uber.org/zap # logger
	go get go.uber.org/zap/zapcore # logger
	go get github.com/natefinch/lumberjack # logger config (for Zap)
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
	go get github.com/confluentinc/confluent-kafka-go/v2/kafka # kafka


generate:
	make generate-mail-api

generate-mail-api:
	mkdir -p pkg/mail_v1
	protoc --proto_path api/mail_v1 --proto_path vendor.protogen \
    --go_out=pkg/mail_v1 --go_opt=paths=source_relative \
    --plugin=protoc-get-go=bin/protoc-gen-go \
    --go-grpc_out=pkg/mail_v1 --go-grpc_opt=paths=source_relative \
    --plugin=protoc-get-go-grpc=bin/protoc-gen-go-grpc \
    api/mail_v1/mail.proto


# mocks
mocks:
	rm -rf tests/mocks && mkdir -p tests/mocks
	#${LOCAL_BIN}/minimock -i github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/repository.IRepository -o tests/unit/mocks -s "_minimock.go"
	#${LOCAL_BIN}/minimock -i github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/service.IService -o tests/unit/mocks -s "_minimock.go"
	#${LOCAL_BIN}/minimock -i github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/utils/auth_helper.AuthHelper -o tests/unit/mocks -s "_minimock.go"

test:
	go clean -testcache
	go test ./... -covermode count -coverpkg=github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/tests/... -count 5
	rm -r tests/unit/controller/logs
	rm -r tests/unit/service/logs

coverage:
	go clean -testcache
	go test ./tests/... -coverprofile=coverage.tmp.out -covermode count  -coverpkg=github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/... -count 5
	grep -v 'mocks\|config' coverage.tmp.out > coverage.out
	rm coverage.tmp.out
	go tool cover -html=coverage.out -o coverage.html;
	go tool cover -func=./coverage.out | grep "total";
	grep -sqFx "/coverage.out" .gitignore || echo "/coverage.out" >> .gitignore
	rm -r tests/unit/controller/logs
	rm -r tests/unit/service/logs
