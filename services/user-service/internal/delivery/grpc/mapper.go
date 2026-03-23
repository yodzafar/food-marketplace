package grpcdelivery

import (
	pb "github.com/yodzafar/food-marketpalce/gen/proto/user"
	"github.com/yodzafar/food-marketpalce/user-service/internal/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func toProto(user *domain.User) *pb.User {
	return &pb.User{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		Role:      string(user.Role),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}

func toGRPCError(err error) error {
	switch err {
	case domain.ErrNotFound:
		return status.Error(codes.NotFound, err.Error())
	case domain.ErrAlreadyExists:
		return status.Error(codes.AlreadyExists, err.Error())
	case domain.ErrUnauthorized:
		return status.Error(codes.Unauthenticated, err.Error())
	case domain.ErrInvalidInput:
		return status.Error(codes.InvalidArgument, err.Error())
	default:
		return status.Error(codes.Internal, "internal server errors")
	}
}
