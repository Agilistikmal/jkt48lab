package model

type Member struct {
	Username  string `json:"username"`
	FullName  string `json:"full_name"`
	Followers int    `json:"followers"`
}
