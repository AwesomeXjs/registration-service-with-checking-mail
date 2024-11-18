compose:
	docker-compose up -d


run-auth:
	cd auth-service && go run cmd/auth/main.go

run-auth-gw:
	cd api-gateway-auth && go run cmd/auth_gw/main.go