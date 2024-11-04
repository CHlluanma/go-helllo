package db

import "testing"

func TestCreate(t *testing.T) {
	Init()
	u := &User{
		UserName: "test",
		Password: "test",
	}
	id, err := CreateUser(u)
	if err != nil {
		t.Log(false)
		return
	}
	t.Log(id)
}
