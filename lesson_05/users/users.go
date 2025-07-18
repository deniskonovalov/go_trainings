package users

import (
	"fmt"
	"github.com/google/uuid"
	"learningGo/lesson_05/documentstore"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Service struct {
	coll documentstore.Collection
}

func NewService(s documentstore.Store) *Service {
	coll, err := s.CreateCollection("users", &documentstore.CollectionConfig{PrimaryKey: "id"})

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &Service{
		coll: *coll,
	}
}

func (s *Service) CreateUser(username string) (*User, error) {
	u := User{
		ID:   uuid.New().String(),
		Name: username,
	}

	doc, err := documentstore.MarshalDocument(u)

	if err != nil {
		return nil, err
	}

	if err = s.coll.Put(*doc); err != nil {
		return nil, err
	}

	return &u, nil
}

func (s *Service) ListUsers() ([]User, error) {
	list := s.coll.List()

	userList := make([]User, 0, len(list))

	for _, doc := range list {
		u := User{}

		if err := documentstore.UnmarshalDocument(&doc, &u); err != nil {
			return nil, err
		}
		userList = append(userList, u)
	}

	return userList, nil
}

func (s *Service) GetUser(userID string) (*User, error) {
	doc, err := s.coll.Get(userID)
	if err != nil {
		return nil, ErrUserNotFound
	}

	u := User{}
	if err := documentstore.UnmarshalDocument(doc, &u); err != nil {
		return nil, err
	}

	return &u, nil
}

func (s *Service) DeleteUser(userID string) error {
	if err := s.coll.Delete(userID); err != nil {
		return ErrUserNotFound
	}

	return nil
}
