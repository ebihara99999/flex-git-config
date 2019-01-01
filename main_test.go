package main

import (
	"os/exec"
	"strings"
	"testing"
	"unsafe"
)

type FakeCommander struct {
	Commands
}

func (c *FakeCommander) GetRepositoriesByGhqList() ([]byte, error) {
	dirs := "/Users/FakeUser/src/github.com/ebihara99999/sample1\n/Users/FakeUser/src/github.com/ebihara99999/sample2"

	out := *(*[]byte)(unsafe.Pointer(&dirs))
	return out, nil
}

func (c *FakeCommander) GetGhqRoot() ([]byte, error) {
	root := "/Users/FakeUser/src"

	out := *(*[]byte)(unsafe.Pointer(&root))
	return out, nil
}

func (c *FakeCommander) ChangeGitUsernameLocally() ([]byte, error) {
	return nil, nil
}

func (c *FakeCommander) ChangeGitEmailLocally(email string) ([]byte, error) {
	return nil, nil
}

func TestMain(t *testing.T) {
	const expected = "Successfully change git config"
	cmd := exec.Command("go", "run", "main.go", "-e", "dev.biibiebi@gmail.com", "-u", "ebihara99999", "-d", "github.com")

	out, _ := cmd.CombinedOutput()
	if strings.Compare(string(out), expected) != 1 {
		t.Errorf("Wanted %s, got %s", expected, string(out))
	}
}
