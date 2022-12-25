package client

import (
	"MongoRedisGin/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
)

type SignUpUserClient struct {
	service pb.AuthServiceClient
}

func NewSignUpUserClient(conn *grpc.ClientConn) *SignUpUserClient {
	service := pb.NewAuthServiceClient(conn)
	return &SignUpUserClient{service}
}
func (signUpUserClient *SignUpUserClient) SignUpUser(credentials *pb.SignUpUserInput) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*50000))
	defer cancel()

	res, err := signUpUserClient.service.SignUpUser(ctx, credentials)

	if err != nil {
		log.Fatalf("SignUpUser: %v", err)
	}

	fmt.Println(res)
}
