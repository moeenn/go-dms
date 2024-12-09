package users

import (
	"context"
	"database/sql"
	"dms/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

// TODO: implement the mock
type IUserRepository interface {
	ListAll(ctx context.Context) (models.UserSlice, error)
	Create(ctx context.Context, input *UserCreateInput) (models.User, error)
}

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) ListAll(ctx context.Context) (models.UserSlice, error) {
	return models.Users().All(ctx, r.DB)
}

// TODO: confirm if this struct is autogenerated somewhere
type UserCreateInput struct {
	Email    string
	Password string
}

func (r *UserRepository) Create(ctx context.Context, input *UserCreateInput) (models.User, error) {
	var newUser models.User
	newUser.Email = input.Email
	newUser.Password = input.Password

	if err := newUser.Insert(ctx, r.DB, boil.Infer()); err != nil {
		return newUser, err
	}

	return newUser, nil
}
