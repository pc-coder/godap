package handler

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
		return h.usersProvider.SearchForUserSearchAttribute(req.FilterValue)
	}}
}

func NewRequestHandlers(usersProvider provider.Users) Handlers {
	return handlers{
		usersProvider: usersProvider,
	}
}
