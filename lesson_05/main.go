package main

import (
	"fmt"
	"learningGo/lesson_05/documentstore"
	"learningGo/lesson_05/users"
	"math/rand"
)

func main() {
	store := documentstore.NewStore()

	service := users.NewService(*store)

	_, err := service.CreateUser("Mike")
	if err != nil {
		fmt.Println(err)
		return
	}

	if _, err = service.CreateUser("John"); err != nil {
		fmt.Println(err)
		return
	}
	if _, err = service.CreateUser("Julie"); err != nil {
		fmt.Println(err)
		return
	}
	if _, err = service.CreateUser("Felix"); err != nil {
		fmt.Println(err)
		return
	}

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
