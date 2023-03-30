package boltdb

import (
	"testing"
	"toolbox/internal/utils"
)

func TestEncodeDecode(t *testing.T) {
	type User struct {
		Name     string
		Password string
		Code     []byte
	}
	users := []*User{
		{
			Name:     "user1aaa",
			Password: "pass1",
		},
		{
			Name:     "user2",
			Password: "pass2",
		},
		{
			Name:     "user3",
			Password: "pass3",
		},
	}

	for _, u := range users {
		data, err := utils.Encode(u)
		if err != nil {
			t.Fatal(err)
		}
		u.Code = data
	}

	for _, u := range users {
		user := new(User)
		err := utils.Decode(u.Code, user)
		if err != nil {
			t.Fatal(err)
		}
		if user.Name != u.Name || user.Password != u.Password {
			t.Error("user decode invalid")
		}
	}
}
