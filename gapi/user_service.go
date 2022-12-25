package gapi

import (
	"MongoRedisGin/config"
	"MongoRedisGin/pb"
	"MongoRedisGin/services"
	"go.mongodb.org/mongo-driver/mongo"
	"html/template"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	config         config.Config
	userService    services.UserService
	userCollection *mongo.Collection
	temp           *template.Template
}

func NewGrpcUserServer(config config.Config, userService services.UserService, userCollection *mongo.Collection, temp *template.Template) (*UserServer, error) {
	userServer := &UserServer{
		config:         config,
		userService:    userService,
		userCollection: userCollection,
		temp:           temp,
	}

	return userServer, nil
}
