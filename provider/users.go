package provider

import (
	"godap/config"
	"godap/godap"
	"strings"
)

type Users interface {
	AreValidCredentials(string, string) bool
	SearchForUserSearchAttribute(filter string) []*godap.LDAPSimpleSearchResultEntry
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

func (u users) SearchForUserSearchAttribute(filter string) []*godap.LDAPSimpleSearchResultEntry {
	ret := make([]*godap.LDAPSimpleSearchResultEntry, 0, 1)

	for _, user := range u.users {
		if strings.Contains(filter, user[u.config.UserSearchAttribute].(string)) {
			ret = append(ret, &godap.LDAPSimpleSearchResultEntry{
				DN:    user["dn"].(string),
				Attrs: user,
			})
		}
	}
	return ret
}

func NewUsersProvider(config config.Data, usersDb []map[string]interface{}) Users {
	return users{
		config: config,
		users:  usersDb,
	}
}
