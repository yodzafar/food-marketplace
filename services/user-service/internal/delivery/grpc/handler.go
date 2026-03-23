package grpcdelivery

import (
	"context"

	pb "github.com/yodzafar/food-marketpalce/gen/proto/user"
	"github.com/yodzafar/food-marketpalce/user-service/internal/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	useCase domain.UserService
}

func NewUserHandler(useCase domain.UserService) *UserHandler {
	return &UserHandler{useCase: useCase}
}

func (h *UserHandler) RegisterService(s *grpc.Server) {
	pb.RegisterUserServiceServer(s, h)
}

func (h *UserHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	if req.Name == "" || req.Email == "" || req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "name, email, password are required")
	}

	user, token, err := h.useCase.Register(ctx, domain.RegisterInput{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		return nil, toGRPCError(err)
	}

	return &pb.RegisterResponse{Token: token, User: toProto(user)}, nil
}

func (h *UserHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	if req.Email == "" || req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "email and password are required")
	}

	user, token, err := h.useCase.Login(ctx, domain.LoginInput{})

	if err != nil {
		return nil, toGRPCError(err)
	}

	return &pb.LoginResponse{Token: token, User: toProto(user)}, nil
}

func (h *UserHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	user, err := h.useCase.GetUser(ctx, req.Id)

	if err != nil {
		return nil, toGRPCError(err)
	}

	return &pb.GetUserResponse{User: toProto(user)}, nil
}

func (h *UserHandler) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	user, err := h.useCase.UpdateUser(ctx, domain.UpdateInput{
		ID:    req.Id,
		Name:  req.Name,
		Phone: req.Phone,
		Role:  string(req.Role),
	})

	if err != nil {
		return nil, toGRPCError(err)
	}

	return &pb.UpdateUserResponse{User: toProto(user)}, nil
}
