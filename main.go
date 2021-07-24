package main

import (
	"fmt"
	"godap/config"
	"godap/godap"
	"godap/utils"
	"log"
	"strings"
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

	hs := make([]godap.LDAPRequestHandler, 0)
	// use a LDAPBindFuncHandler to provide a callback function to respond to bind requests
	hs = append(hs, &godap.LDAPBindFuncHandler{LDAPBindFunc: func(binddn string, bindpw []byte) bool {
		for _, val := range users {
			loginAttribute := val[configData.UserLoginAttribute].(string)
			password := val["password"]
			if strings.Contains(binddn, loginAttribute) && string(bindpw) == password {
				return true
			}
		}
		return false
	}})

	// use a LDAPSimpleSearchFuncHandler to reply to search queries
	hs = append(hs, &godap.LDAPSimpleSearchFuncHandler{LDAPSimpleSearchFunc: func(req *godap.LDAPSimpleSearchRequest) []*godap.LDAPSimpleSearchResultEntry {

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

	}})

	s := &godap.LDAPServer{
		Handlers: hs,
	}

	fmt.Println("Starting mock LDAP server on ", configData.Port)
	err = s.ListenAndServe(fmt.Sprintf(":%d", configData.Port))
	if err != nil {
		fmt.Printf("Failed to start server. Error : %v", err)
	}
}
