package usecase

import (
	"context"
	"fmt"

	"github.com/yodzafar/food-marketpalce/user-service/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	repo  domain.UserRepository
	cache domain.UserCache
}

func NewUserUseCase(repo domain.UserRepository, cache domain.UserCache) domain.UserService {
	return &userUseCase{repo: repo, cache: cache}
}

func (u *userUseCase) Register(ctx context.Context, req domain.RegisterInput) (*domain.User, string, error) {
	//TODO implement me
	_, err := u.repo.GetByEmail(ctx, req.Email)
	if err == nil {
		return nil, "", domain.ErrAlreadyExists
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, "", err
	}

	user := &domain.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: string(hash),
	}

	if err := u.repo.Create(ctx, user); err != nil {
		return nil, "", fmt.Errorf("create user: %w", err)
	}

	token, err := generateToken(user)

	if err != nil {
		return nil, "", fmt.Errorf("generate token: %w", err)
	}

	_ = u.cache.Set(ctx, user)

	return user, token, nil
}

func (u *userUseCase) Login(ctx context.Context, req domain.LoginInput) (*domain.User, string, error) {
	user, err := u.repo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, "", domain.ErrCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, "", domain.ErrCredentials
	}

	token, err := generateToken(user)

	if err != nil {
		return nil, "", fmt.Errorf("generate token: %w", err)
	}

	_ = u.cache.Set(ctx, user)

	return user, token, nil
}

func (u *userUseCase) GetUser(ctx context.Context, id string) (*domain.User, error) {
	user, err := u.cache.Get(ctx, id)

	if err == nil {
		return user, nil
	}

	user, err = u.repo.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}

	_ = u.cache.Set(ctx, user)

	return user, nil
}

func (u *userUseCase) UpdateUser(ctx context.Context, req domain.UpdateInput) (*domain.User, error) {
	user, err := u.repo.GetByID(ctx, req.ID)

	if err != nil {
		return nil, err
	}

	user.Name = req.Name
	user.Phone = req.Phone

	if err := u.repo.Update(ctx, user); err != nil {
		return nil, fmt.Errorf("update user: %w", err)
	}

	_ = u.cache.Delete(ctx, req.ID)

	return user, nil
}
