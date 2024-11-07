compose:
	docker-compose up -d


run-auth:
	cd server/auth-service && go run cmd/auth/main.go

run-auth-gw:
	cd server/api-gateway-auth && go run cmd/auth_gw/main.go