package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/yodzafar/food-marketpalce/user-service/internal/domain"
)

type userPostgresRepo struct {
	db *sqlx.DB
	sq sq.StatementBuilderType
}

func NewUserRepository(db *sqlx.DB) domain.UserRepository {
	return &userPostgresRepo{
		db: db,
		sq: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (ur userPostgresRepo) Create(ctx context.Context, user *domain.User) error {
	user.ID = uuid.NewString()

	query, args, err := ur.sq.
		Insert("users").
		Columns("id", "name", "email", "password_hash", "phone", "role").
		Values(user.ID, user.Name, user.Email, user.PasswordHash, user.Phone, user.Role).
		ToSql()

	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	_, err = ur.db.ExecContext(ctx, query, args...)

	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (ur userPostgresRepo) GetByID(ctx context.Context, id string) (*domain.User, error) {
	query, args, err := ur.sq.
		Select("id", "name", "email", "phone", "role", "created_at", "updated_at").
		From("users").
		Where(sq.Eq{"id": id}).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}

	user := &domain.User{}
	err = ur.db.QueryRowxContext(ctx, query, args...).StructScan(user)

	if !errors.Is(err, sql.ErrNoRows) {
		return nil, domain.ErrNotFound
	}

	if err != nil {
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}

	return user, nil
}

func (ur userPostgresRepo) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (ur userPostgresRepo) Update(ctx context.Context, user *domain.User) error {
	//TODO implement me
	panic("implement me")
}
