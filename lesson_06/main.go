package main

import (
	"fmt"
	"learningGo/lesson_06/documentstore"
	"learningGo/lesson_06/users"
)

func main() {
	s := documentstore.NewStore()

	service := users.NewService(*s)

	usernames := [6]string{
		"Mike",
		"John",
		"Julie",
		"Felix",
		"Rossy",
		"Dan",
	}

	for _, username := range usernames {
		if _, err := service.CreateUser(username); err != nil {
			fmt.Println(err)
			return
		}
	}

	dump, err := s.Dump()

	if err != nil {
		fmt.Println(err)
		return
	}

	if _, err := service.CreateUser("New User for only this store"); err != nil {
		fmt.Println(err)
		return
	}

	restored, err := documentstore.NewStoreFromDump(dump)
	if err != nil {
		fmt.Println(err)
		return
	}

	serviceFromRestored := users.NewService(*restored)

	list, err := service.ListUsers()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Original store:")
	fmt.Printf("%+v\n", list)

	listFromDump, err := serviceFromRestored.ListUsers()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Restored store:")
	fmt.Printf("%+v\n", listFromDump)

	if err := s.DumpToFile("users.dump"); err != nil {
		fmt.Println(err)
		return
	}

	newStore, err := documentstore.NewStoreFromFile("users.dump")
	if err != nil {
		fmt.Println(err)
		return
	}
	serviceFromFile := users.NewService(*newStore)

	usersFromFile, err := serviceFromFile.ListUsers()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Users from file:")
	fmt.Printf("%+v\n", usersFromFile)
}
