package controller

import (
	"app/models"
)

var Users []models.User

func CreateUser(data models.User) {
	Users = append(Users, data)
}

func GetListUser() []models.User {
	return Users
}

// getbyid
func GetUserById(id int) (models.User, bool) {
	for _, val := range Users {
		if val.Id == id {
			return val, false
		}
	}

	return models.User{}, true
}

// update
func UpdateUser(id int) (models.User, bool) {
	for i, val := range Users {
		if val.Id == id {
			Users[i].Name = "changed Name"
			Users[i].Surname = "changed Surname"
			return Users[i], false
		}
	}

	return models.User{}, true
}

// delete
func DeleteUser(id int) bool {
	for i, val := range Users {
		if val.Id == id {
			Users = append(Users[:i], Users[i+1:]...)
			return false
		}
	}

	return true
}
