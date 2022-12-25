package gapi

import (
	"MongoRedisGin/config"
	"MongoRedisGin/pb"
	"MongoRedisGin/services"
	"go.mongodb.org/mongo-driver/mongo"
	"html/template"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	config.Config
	authService    services.AuthService
	userService    services.UserService
	userCollection *mongo.Collection
	temp           *template.Template
}

func NewGrpcAuthServer(config config.Config, authService services.AuthService, userService services.UserService, userCollection *mongo.Collection, temp *template.Template) (*AuthServer, error) {
	return &AuthServer{
		Config:         config,
		authService:    authService,
		userService:    userService,
		userCollection: userCollection,
		temp:           temp,
	}, nil

}
