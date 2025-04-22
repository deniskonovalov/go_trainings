package users

import (
	"github.com/stretchr/testify/assert"
	"lesson_07/internal/documentstore"
	"testing"
)

func Test_NewService(t *testing.T) {
	s := documentstore.NewStore()

	service := NewService(s)

	assert.NotNil(t, service)
	assert.IsType(t, &Service{}, service)
	assert.IsType(t, documentstore.Collection{}, service.coll)
}

func Test_NewService_WithExistingCollection(t *testing.T) {
	s := documentstore.NewStore()

	_, err := s.CreateCollection("users", &documentstore.CollectionConfig{PrimaryKey: "id"})
	assert.NoError(t, err)

	service := NewService(s)

	assert.NotNil(t, service)
	assert.IsType(t, &Service{}, service)
	assert.IsType(t, documentstore.Collection{}, service.coll)
}

func Test_CreateUser(t *testing.T) {
	s := documentstore.NewStore()
	service := NewService(s)

	u, err := service.CreateUser("Mike")
	assert.NoError(t, err)

	assert.NotNil(t, u)
	assert.IsType(t, User{}, *u)
	assert.Equal(t, "Mike", u.Name)
	assert.NotEmpty(t, u.ID)
}

func Test_ListUsers(t *testing.T) {
	s := documentstore.NewStore()
	service := NewService(s)
	user1, err := service.CreateUser("Mike")
	assert.NoError(t, err)
	user2, err := service.CreateUser("John")
	assert.NoError(t, err)
	user3, err := service.CreateUser("Julie")
	assert.NoError(t, err)

	l, err := service.ListUsers()
	assert.NoError(t, err)
	assert.Len(t, l, 3)
	assert.IsType(t, []User{}, l)

	assert.Contains(t, l, *user1)
	assert.Contains(t, l, *user2)
	assert.Contains(t, l, *user3)
	assert.ObjectsAreEqual(l[0], *user1)
	assert.ObjectsAreEqual(l[1], *user2)
	assert.ObjectsAreEqual(l[2], *user3)
}

func Test_GetUser(t *testing.T) {
	s := documentstore.NewStore()

	service := NewService(s)
	user, err := service.CreateUser("Mike")

	assert.NoError(t, err)
	u, err := service.GetUser(user.ID)

	assert.NoError(t, err)
	assert.Equal(t, user.ID, u.ID)
	assert.Equal(t, user.Name, u.Name)
}

func Test_GetUser_NotFound(t *testing.T) {
	s := documentstore.NewStore()

	service := NewService(s)
	u, err := service.GetUser("non-existing-id")

	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrUserNotFound)
	assert.Nil(t, u)
}

func Test_DeleteUser(t *testing.T) {
	s := documentstore.NewStore()
	service := NewService(s)
	user, err := service.CreateUser("Mike")
	assert.NoError(t, err)

	err = service.DeleteUser(user.ID)

	assert.NoError(t, err)

	u, err := service.GetUser(user.ID)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrUserNotFound)
	assert.Nil(t, u)
}

func Test_DeleteUser_NotFound(t *testing.T) {
	s := documentstore.NewStore()
	service := NewService(s)

	err := service.DeleteUser("non-existing-id")
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrUserNotFound)
}
