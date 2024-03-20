package gapi

import (
	db "github.com/dpurbosakti/booknest-grpc/internal/db/sqlc"
	"github.com/dpurbosakti/booknest-grpc/internal/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *pb.User {
	return &pb.User{
		Name:              user.Name,
		Phone:             user.Phone,
		Email:             user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
		UpdatedAt:         timestamppb.New(user.UpdatedAt),
	}
}
