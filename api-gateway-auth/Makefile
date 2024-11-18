LOCAL_BIN:=$(CURDIR)/bin

run:
	go run cmd/auth_gw/main.go

debug:
	go run cmd/auth_gw/main.go -l debug

lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml

swagger:
	$(LOCAL_BIN)/swag init -g cmd/auth_gw/main.go


install-deps:
	mkdir -p $(LOCAL_BIN)
	GOBIN=$(LOCAL_BIN) go install github.com/gojuno/minimock/v3/cmd/minimock@v3.4.0 # mocks
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.60.3 # lint
	GOBIN=$(LOCAL_BIN) go install github.com/swaggo/swag/cmd/swag@latest # swagger


get-deps:
	go get -u github.com/brianvoe/gofakeit # fake data
	go get github.com/joho/godotenv # env config
	go get github.com/pkg/errors # errors
	go get github.com/stretchr/testify/require # tests
	go get github.com/golang/mock/gomock # mocks
	go get github.com/rs/cors # cors
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
	go get github.com/goccy/go-json # json decoder


# mocks
mocks:
	rm -rf tests/unit/mocks && mkdir -p tests/unit/mocks
	${LOCAL_BIN}/minimock -i github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/client/auth_client.AuthClient -o tests/unit/mocks -s "_minimock.go"
	${LOCAL_BIN}/minimock -i github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/utils/header_helper.IHeaderHelper -o tests/unit/mocks -s "_minimock.go"

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
