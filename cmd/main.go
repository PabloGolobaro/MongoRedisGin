package main

import (
	"MongoRedisGin/client"
	"MongoRedisGin/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const (
	address = "localhost:8080"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())

	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	defer conn.Close()

	// Sign Up
	if false {
		signUpUserClient := client.NewSignUpUserClient(conn)
		newUser := &pb.SignUpUserInput{
			Name:            "Lera Golovina",
			Email:           "LeraGolo@gmail.com",
			Password:        "password123",
			PasswordConfirm: "password123",
		}
		signUpUserClient.SignUpUser(newUser)
	}

	// Sign In
	if false {
		signInUserClient := client.NewSignInUserClient(conn)

		credentials := &pb.SignInUserInput{
			Email:    "leragolo@gmail.com",
			Password: "password123",
		}
		signInUserClient.SignInUser(credentials)
	}

	// Get Me
	if true {

		getMeClient := client.NewGetMeClient(conn)
		id := &pb.GetMeRequest{
			Id: "63a8673c818c5918879860a0",
		}
		getMeClient.GetMeUser(id)

	}

}
