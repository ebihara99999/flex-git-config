package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

var name = flag.String("name", "World", "A name to say hello to")

var username string
var email string
var domain string
var noop bool

func init() {
	flag.StringVar(&username, "username", "", "Set Github/GHE username")
	flag.StringVar(&username, "u", "", "Set Github/GHE username")
	flag.StringVar(&email, "email", "", "Set Github/GHE email")
	flag.StringVar(&email, "e", "", "Set Github/GHE email")
	flag.StringVar(&domain, "domain", "", "Set Github/GHE domain")
	flag.StringVar(&domain, "d", "", "Set Github/GHE domain")
	flag.BoolVar(&noop, "n", false, "noop")

}

func main() {
	flag.Parse()

	// Check the args
	usage := "Usage: hoge -u username -e email -d domain"
	if username == "" {
		panic(usage)
	}

	if email == "" {
		panic(usage)
	}

	if domain == "" {
		panic(usage)
	}

	// directoris gotten by `ghq list`
	repositories := []string{}
	_tmpRepositories, _ := exec.Command("ghq", "list").Output()
	repositories = strings.Split(string(_tmpRepositories), "\n")

	// ghq root path
	_ghqRoot, _ := exec.Command("ghq", "root").Output()
	ghqRoot := string(_ghqRoot)
	ghqRoot = strings.TrimRight(ghqRoot, "\n")

	// directories matched the domain specified by the d option
	matchedRepositories := []string{}
	r := regexp.MustCompile(domain)
	for _, repo := range repositories {
		if r.MatchString(repo) {
			_absPath := path.Join(ghqRoot, repo)
			matchedRepositories = append(matchedRepositories, _absPath)
		}
	}

	// For debug
	if noop {
		fmt.Println("Noop operation.")
		fmt.Println("The taget repositories are:")
		for _, targetDir := range matchedRepositories {
			fmt.Println(targetDir)
		}
		os.Exit(0)
	}

	prevDir, _ := filepath.Abs(".")
	defer os.Chdir(prevDir)

	for _, targetDir := range matchedRepositories {
		err := os.Chdir(targetDir)
		if err != nil {
			fmt.Println("Error occurs when changing directory")
			panic(err)
		}

		errUsername := exec.Command("git", "config", "--local", "user.name", username).Run()

		if errUsername != nil {
			fmt.Println("Error occurs when executing `git config --local user.name`")
			panic(errUsername)
		}

		errEmail := exec.Command("git", "config", "--local", "user.email", email).Run()

		if errEmail != nil {
			fmt.Println("Error occurs when executing `git config --local user.email`")
			panic(errEmail)
		}
	}
}
