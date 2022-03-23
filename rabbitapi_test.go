package main

import "testing"

func TestCreateNewUser(t *testing.T) {
	if _, err := ReqCreateNewUser("ampas", "ampas"); err != nil {
		t.Fatal("error: ", err)
	}
}
