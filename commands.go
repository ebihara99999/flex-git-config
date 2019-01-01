package main

import (
	"fmt"
	"os/exec"
)

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
	fmt.Println("hogehogehogehogehogehogeho")
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
