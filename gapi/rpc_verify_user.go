package gapi

import (
	"MongoRedisGin/pb"
	"MongoRedisGin/utils"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *AuthServer) VerifyEmail(ctx context.Context, req *pb.VerifyEmailRequest) (*pb.GenericResponce, error) {
	code := req.GetVerificationCode()
	verificationCode := utils.Encode(code)
	query := bson.D{{Key: "verificationCode", Value: verificationCode}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "verified", Value: true}}}}
	result, err := server.userCollection.UpdateOne(ctx, query, update)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	if result.MatchedCount == 0 {
		return nil, status.Errorf(codes.PermissionDenied, err.Error())
	}
	res := &pb.GenericResponce{
		Status:  "Success",
		Message: "Email verified successfully",
	}
	return res, nil
}
