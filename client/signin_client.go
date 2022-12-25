package client

import (
	"MongoRedisGin/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
)

type SignInUserClient struct {
	service pb.AuthServiceClient
}

func NewSignInUserClient(conn *grpc.ClientConn) *SignInUserClient {
	service := pb.NewAuthServiceClient(conn)

	return &SignInUserClient{service}
}

func (signInUserClient *SignInUserClient) SignInUser(credentials *pb.SignInUserInput) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	res, err := signInUserClient.service.SignInUser(ctx, credentials)

	if err != nil {
		log.Fatalf("SignInUser: %v", err)
	}

	fmt.Println(res)
}
