package main

import (
	"fmt"

	"app/controller"
	"app/models"
)

func main() {

	controller.CreateUser(
		models.User{
			Id:      1,
			Name:    "Shohruh",
			Surname: "Safarov",
		},
	)

	controller.CreateUser(
		models.User{
			Id:      2,
			Name:    "Abduqodir",
			Surname: "Musayev",
		},
	)

	for _, user := range controller.GetListUser() {
		fmt.Println(user)
	}
	// Get By Id
	user, err := controller.GetUserById(3)
	fmt.Println("Get By id:", user, "error:", err)
	// Update User
	userUpdated, err2 := controller.UpdateUser(2)
	fmt.Println("Updated user:", userUpdated, "error:", err2)
	fmt.Println(controller.GetListUser())
	// Delete User
	err3 := controller.DeleteUser(2)
	fmt.Println("error for deleting:", err3)
	fmt.Println(controller.GetListUser())

}
