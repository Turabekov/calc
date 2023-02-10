package main

import (
	"fmt"

	"app/controller"
	"app/models"
)

func main() {

	controller.GenerateUser(100)

	fmt.Println("Page raqamini kiriting:")
	pageNumber := 0
	fmt.Scan(&pageNumber)

	pageOffset := pageNumber * 10

	users, err := controller.GetListUser(models.GetListRequest{
		Offset: pageOffset,
		Limit:  10,
	})

	if err {
		fmt.Println("users out of range")
		return
	}

	for _, user := range users {
		fmt.Println(user)
	}
}
