package main

import (
	"fmt"

	"github.com/codersgyan/podcast/user"
	"github.com/fatih/color"

	"github.com/codersgyan/podcast/auth"
)

func main() {
	auth.LoginWithCredentials("ashu", "secret")
	session := auth.GetSession()
	fmt.Println("session", session)

	user := user.User{
		Email: "user@gmail.com",
		Name:  "John doe",
	}

	// fmt.Println(user.Email, user.Name)
	color.Green(user.Email)
}
