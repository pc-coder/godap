package config

type Data struct {
	Port                int    `json:"port"`
	UserLoginAttribute  string `json:"userLoginAttribute"`
	UserSearchAttribute string `json:"userSearchAttribute"`
}
