package users

import (
	"context"
	"dms/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type UserRepository struct {
	DB boil.ContextExecutor
}

func NewUserRepository(db boil.ContextExecutor) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) ListAll(ctx context.Context) (models.UserSlice, error) {
	return models.Users().All(ctx, r.DB)
}

type UserInput struct {
	Email    string
	Name     string
	Password string
	Role     models.UserRole
}

func (r *UserRepository) Create(ctx context.Context, input *UserInput) (models.User, error) {
	var newUser models.User
	newUser.Email = input.Email
	newUser.Name = input.Name
	newUser.Password = input.Password
	newUser.Role = input.Role

	if err := newUser.Insert(ctx, r.DB, boil.Infer()); err != nil {
		return newUser, err
	}

	return newUser, nil
}
