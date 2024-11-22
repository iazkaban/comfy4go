package client

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type User struct {
	UserID   string
	UserName string
}

type UserList struct {
	Storage string            `json:"storage"`
	M       map[string]string `json:"users"`
	Users   []*User           `json:"-"`
}

func (client *Client) GetUsers() (*UserList, error) {
	apiUri := "/api/users"
	apiMethod := http.MethodGet

	apiUrl := url.URL{
		Scheme: "http",
		Host:   client.host,
		Path:   apiUri,
	}

	req, err := http.NewRequest(apiMethod, apiUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	userList := &UserList{
		Storage: "",
		M:       make(map[string]string),
	}

	err = json.Unmarshal(body, userList)
	if err != nil {
		panic(err)
		return nil, err
	}

	if len(userList.M) > 0 {
		userList.Users = make([]*User, 0, len(userList.M))
		for k, v := range userList.M {
			u := &User{
				UserID:   k,
				UserName: v,
			}
			userList.Users = append(userList.Users, u)
		}
		userList.M = nil
	}

	return userList, nil
}

func (client *Client) SelectUser(userID string) {
	client.UserID = userID
}
