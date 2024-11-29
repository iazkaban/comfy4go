package model

type User struct {
	UserID   string
	UserName string
}

type UserList struct {
	Storage string            `json:"storage"`
	M       map[string]string `json:"users"`
	Users   []*User           `json:"-"`
}
