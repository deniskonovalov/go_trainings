package main

import (
	"fmt"
	"learningGo/lesson_05/documentstore"
	"learningGo/lesson_05/users"
	"math/rand"
)

func main() {
	store := documentstore.NewStore()

	coll, err := store.CreateCollection("users", &documentstore.CollectionConfig{PrimaryKey: "id"})
	if err != nil {
		fmt.Println(err)
		return
	}

	service := users.NewService(*coll)

	_, err = service.CreateUser("Mike")
	if err != nil {
		fmt.Println(err)
		return
	}

	_, _ = service.CreateUser("John")
	_, _ = service.CreateUser("Julie")
	_, _ = service.CreateUser("Felix")

	l, err := service.ListUsers()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%#v\n", l)

	randUser := l[rand.Intn(len(l)-1)]

	userFromStorage, err := service.GetUser(randUser.ID)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(userFromStorage)

	err = service.DeleteUser(randUser.ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	l, err = service.ListUsers()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%#v\n", l)
}
