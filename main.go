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
	"os/exec"
)

var username string
var email string
var domain string
var noop bool

func init() {
	flag.StringVar(&username, "username", "", "Set the username used in the service domain pointed by -domain option")
	flag.StringVar(&username, "u", "", "Set the username used in the service domain pointed by -domain option")
	flag.StringVar(&email, "email", "", "Set the email used in the service domain pointed by -domain option")
	flag.StringVar(&email, "e", "", "Set the email used in the service domain pointed by -domain option")
	flag.StringVar(&domain, "domain", "", "Set a service domain like github.com or bitbucket.org. A self-hosted custom domain is also supported.")
	flag.StringVar(&domain, "d", "", "Set a service domain like github.com or bitbucket.org. A self-hosted custom domain is also supported.")
	flag.BoolVar(&noop, "n", false, "Noop option. Show the repositories that will be affected.")
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

// Commands interface is
type Commands interface {
	GetRepositoriesByGhqList() ([]byte, error)
	GetGhqRoot() ([]byte, error)
	ChangeGitUsernameLocally(username string) ([]byte, error)
	ChangeGitEmailLocally(email string) ([]byte, error)
}

// Commander is 
type Commander struct {
	c Commands
}

// GetRepositoriesByGhqList is
func (c *Commander) GetRepositoriesByGhqList() ([]byte, error) {
	cmd := exec.Command("ghq", "list")
	out, err := cmd.CombinedOutput()
	return out, err
}

// GetGhqRoot is
func (c *Commander) GetGhqRoot() ([]byte, error) {
	cmd := exec.Command("ghq", "root")
	out, err := cmd.CombinedOutput()
	return out, err
}

// ChangeGitUsernameLocally is
func (c *Commander) ChangeGitUsernameLocally(username string) ([]byte, error) {
	cmd := exec.Command("git", "config", "--local", "user.name", username)
	out, err := cmd.CombinedOutput()
	return out, err
}

// ChangeGitEmailLocally is
func (c *Commander) ChangeGitEmailLocally(email string) ([]byte, error) {
	cmd := exec.Command("git", "config", "--local", "user.email", email)
	out, err := cmd.CombinedOutput()
	return out, err
}
