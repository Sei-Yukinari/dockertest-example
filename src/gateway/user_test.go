package gateway

import (
	"context"
	"dockertest-example/src/domain/model"
	"dockertest-example/src/test"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setup(t *testing.T) (context.Context, *gorm.DB, *User) {
	resource, pool := test.CreateMySQLContainer("../../db/schema.sql")
	ctx := context.Background()
	db := test.ConnectMySQLContainer(resource, pool, t)
	repo := NewUser(db)
	return ctx, db, repo
}

func TestUser_Gateway(t *testing.T) {
	t.Parallel()
	ctx, db, repo := setup(t)
	t.Run("Get User", func(t *testing.T) {
		actual := &model.User{
			ID:   2,
			Name: "Dummy Name",
		}

		test.Seeds(db,
			[]interface{}{
				actual,
			})

		res, err := repo.FindByID(ctx, 2)
		assert.Equal(t, err, nil)
		assert.Equal(t, res.ID, actual.ID)
		assert.Equal(t, res.Name, actual.Name)
	})
	t.Run("Create User", func(t *testing.T) {
		actual := &model.User{
			ID:   2,
			Name: "Dummy Name",
		}
		err := repo.Create(ctx, actual)
		assert.NoError(t, err)
	})

}
