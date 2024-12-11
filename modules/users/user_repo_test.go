package users

import (
	"context"
	"database/sql"
	"dms/config"
	"dms/models"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestUserRepo(t *testing.T) {
	cfg, err := config.NewConfig()
	assert.NoError(t, err)

	db, err := sql.Open("postgres", cfg.Database.ConnectionURI)
	assert.NoError(t, err)
	ctx := context.TODO()

	tx, err := db.BeginTx(ctx, nil)
	assert.NoError(t, err)
	userRepo := NewUserRepository(tx)

	input := &UserInput{
		Email:    "sample@site.com",
		Name:     "Admin",
		Password: "123123123",
		Role:     models.UserRoleAdmin,
	}

	t.Run("create new valid user", func(st *testing.T) {
		newUser, err := userRepo.Create(ctx, input)
		assert.NoError(st, err)

		assert.NotNil(st, newUser.UserID)
		assert.Equal(st, input.Email, newUser.Email)
		assert.Equal(st, input.Name, newUser.Name)
		assert.Equal(st, input.Password, newUser.Password)
		assert.Equal(st, input.Role, newUser.Role)

		st.Cleanup(func() {
			_, err := newUser.Delete(ctx, tx)
			assert.NoError(st, err)
		})
	})

	t.Run("list users", func(st *testing.T) {
		newUser, err := userRepo.Create(ctx, input)
		assert.NoError(st, err)

		allUsers, err := userRepo.ListAll(ctx)
		assert.NoError(st, err)
		assert.True(st, len(allUsers) > 0)

		found := false
		for _, user := range allUsers {
			if user.Email == input.Email {
				found = true
				break
			}
		}
		assert.True(st, found)

		st.Cleanup(func() {
			_, err := newUser.Delete(ctx, tx)
			assert.NoError(st, err)
		})
	})
}
