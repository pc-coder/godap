package main

import (
	"godap/config"
	"godap/godap"
	"strings"
)

type Handlers interface {
	GetBindHandler() *godap.LDAPBindFuncHandler
	GetSearchHandler() *godap.LDAPSimpleSearchFuncHandler
}

type handlers struct {
	config config.Data
	users  []map[string]interface{}
}

// GetBindHandler returns a handler function to respond to bind requests
func (h handlers) GetBindHandler() *godap.LDAPBindFuncHandler {
	return &godap.LDAPBindFuncHandler{LDAPBindFunc: func(binddn string, bindpw []byte) bool {
		for _, val := range h.users {
			loginAttribute := val[h.config.UserLoginAttribute].(string)
			password := val["password"]
			if strings.Contains(binddn, loginAttribute) && string(bindpw) == password {
				return true
			}
		}
		return false
	}}
}

// GetSearchHandler returns a handler for simple search requests
func (h handlers) GetSearchHandler() *godap.LDAPSimpleSearchFuncHandler {
	return &godap.LDAPSimpleSearchFuncHandler{LDAPSimpleSearchFunc: func(req *godap.LDAPSimpleSearchRequest) []*godap.LDAPSimpleSearchResultEntry {

		ret := make([]*godap.LDAPSimpleSearchResultEntry, 0, 1)

		// here we produce a single search result that matches whatever
		// they are searching for
		if req.FilterAttr == "uid" {
			ret = append(ret, &godap.LDAPSimpleSearchResultEntry{
				DN: "cn=" + req.FilterValue + "," + req.BaseDN,
				Attrs: map[string]interface{}{
					"sn":            req.FilterValue,
					"cn":            req.FilterValue,
					"uid":           req.FilterValue,
					"homeDirectory": "/home/" + req.FilterValue,
					"objectClass": []string{
						"top",
						"posixAccount",
						"inetOrgPerson",
					},
				},
			})
		}

		return ret

	}}
}

func NewRequestHandlers(config config.Data, users []map[string]interface{}) Handlers {
	return handlers{
		config: config,
		users:  users,
	}
}
