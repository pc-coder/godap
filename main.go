package main

import (
	"fmt"
	"godap/config"
	"godap/godap"
	"godap/handler"
	"godap/provider"
	"godap/utils"
	"log"
)

const (
	ServerConfigPATH = "./config/ldap-server-mock-conf.json"
	UserDatabasePATH = "./config/users.json"
)

func main() {
	var configData config.Data
	err := utils.LoadJSONFile(ServerConfigPATH, &configData)
	if err != nil {
		log.Fatalln(err)
	}

	var users []map[string]interface{}
	err = utils.LoadJSONFile(UserDatabasePATH, &users)
	if err != nil {
		log.Fatalln(err)
	}

	usersProvider := provider.NewUsersProvider(configData, users)
	requestHandlers := handler.NewRequestHandlers(usersProvider)

	hs := []godap.LDAPRequestHandler{requestHandlers.GetBindHandler(), requestHandlers.GetSearchHandler()}

	s := &godap.LDAPServer{
		Handlers: hs,
	}

	fmt.Println("Starting mock LDAP server on ", configData.Port)
	err = s.ListenAndServe(fmt.Sprintf(":%d", configData.Port))
	if err != nil {
		fmt.Printf("Failed to start server. Error : %v", err)
	}
}
