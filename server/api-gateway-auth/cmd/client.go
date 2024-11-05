package main

import (
	"context"
	"fmt"
	"log"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/pkg/auth_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	address = "localhost:50051"
)

func main() {
	ctx := context.Background()

	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := auth_v1.NewAuthV1Client(conn)
	r, err := client.Login(ctx, &auth_v1.LoginRequest{Email: "awesomex@gmail.com", Password: "qwerty123"})
	fmt.Println(r)

}
