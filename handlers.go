package main

import (
	"godap/godap"
	"godap/provider"
)

type Handlers interface {
	GetBindHandler() *godap.LDAPBindFuncHandler
	GetSearchHandler() *godap.LDAPSimpleSearchFuncHandler
}

type handlers struct {
	usersProvider provider.Users
}

// GetBindHandler returns a handler function to respond to bind requests
func (h handlers) GetBindHandler() *godap.LDAPBindFuncHandler {
	return &godap.LDAPBindFuncHandler{LDAPBindFunc: func(binddn string, bindpw []byte) bool {
		return h.usersProvider.AreValidCredentials(binddn, string(bindpw))
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

func NewRequestHandlers() Handlers {
	return handlers{}
}
