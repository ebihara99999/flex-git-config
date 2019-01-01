package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

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
	Run()
}

// Run is the main function
func Run() {
	flag.Parse()

	_, err := checkArgs()
	if err != nil {
		panic(err)
	}

	c := Commander{}

	// directoris gotten by `ghq list`
	out, _ := c.GetRepositoriesByGhqList()
	repositories := strings.Split(string(out), "\n")

	// ghq root path
	_ghqRoot, _ := c.GetGhqRoot()
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

		_, errUsername := c.ChangeGitUsernameLocally(username)

		if errUsername != nil {
			fmt.Println("Error occurs when executing `git config --local user.name`")
			panic(errUsername)
		}

		_, errEmail := c.ChangeGitEmailLocally(email)

		if errEmail != nil {
			fmt.Println("Error occurs when executing `git config --local user.email`")
			panic(errEmail)
		}
	}
	fmt.Println("Successfully change git config")
}

func checkArgs() (result bool, err error) {
	usage := "Usage: flex-git-config -u username -e email -d domain"
	if username == "" {
		return false, errors.New(usage)
	}

	if email == "" {
		return false, errors.New(usage)
	}

	if domain == "" {
		return false, errors.New(usage)
	}
	return true, nil
}


