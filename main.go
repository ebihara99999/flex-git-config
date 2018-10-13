package main

import (
	"flag"
	"os/exec"
)

var name = flag.String("name", "World", "A name to say hello to")

var username string
var email string

func init() {
	flag.StringVar(&username, "username", "", "Set Github/GHE username")
	flag.StringVar(&username, "u", "", "Set Github/GHE username")
	flag.StringVar(&email, "email", "", "Set Github/GHE email")
	flag.StringVar(&email, "e", "", "Set Github/GHE email")
}

func main() {
	flag.Parse()

	usage := "Usage: hoge -u username -e email -"
	if username == "" {
		panic(usage)
	}

	errUsername := exec.Command("git", "config", "--local", "user.name", username).Run()

	if errUsername != nil {
		panic(errUsername)
	}

	if email == "" {
		panic(usage)
	}

	errEmail := exec.Command("git", "config", "--local", "user.email", email).Run()

	if errEmail != nil {
		panic(errEmail)
	}
}
