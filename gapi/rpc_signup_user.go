package gapi

import (
	"MongoRedisGin/models"
	"MongoRedisGin/pb"
	"MongoRedisGin/utils"
	"context"
	"github.com/thanhpk/randstr"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

func (server *AuthServer) SignUpUser(ctx context.Context, req *pb.SignUpUserInput) (*pb.GenericResponce, error) {
	if req.GetPassword() != req.GetPasswordConfirm() {
		return nil, status.Errorf(codes.InvalidArgument, "passwords do not match")
	}
	user := models.SignUpInput{
		Name:            req.GetName(),
		Email:           req.GetEmail(),
		Password:        req.GetPassword(),
		PasswordConfirm: req.GetPasswordConfirm(),
	}

	newUser, err := server.authService.SignUpUser(&user)
	if err != nil {
		if strings.Contains(err.Error(), "email already exist") {
			return nil, status.Errorf(codes.AlreadyExists, "%s", err.Error())

		}
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}
	code := randstr.String(20)
	verificationCode := utils.Encode(code)
	server.userService.UpdateUserById(newUser.ID.Hex(), "verificationCode", verificationCode)

	var firstName = newUser.Name
	if strings.Contains(firstName, " ") {
		firstName = strings.Split(firstName, " ")[0]
	}

	// ? Send Email
	emailData := utils.EmailData{
		URL:       server.Config.Origin + "/verifyemail/" + code,
		FirstName: firstName,
		Subject:   "Your account verification code",
	}
	err = utils.SendEmail(newUser, &emailData, server.temp, "verificationCode.html")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "There was an error: %s", err.Error())
	}
	message := "We sent an email with a verification code to " + newUser.Email

	res := &pb.GenericResponce{
		Status:  "success",
		Message: message,
	}
	return res, nil
}
