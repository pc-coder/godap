package provider

import (
	"godap/config"
	"strings"
)

type Users interface {
	AreValidCredentials(string, string) bool
}

type users struct {
	config config.Data
	users  []map[string]interface{}
}

func (u users) AreValidCredentials(username string, password string) bool {
	for _, val := range u.users {
		userLoginAttribute := val[u.config.UserLoginAttribute].(string)
		userPassword := val["password"]
		if strings.Contains(username, userLoginAttribute) && password == userPassword {
			return true
		}
	}
	return false
}

func NewUsersRepository() Users {
	return users{}
}
