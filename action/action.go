package action

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

type StepInputs map[string]string

func tempDir() string {
	i, err := rand.Int(rand.Reader, big.NewInt(100000000))
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("/tmp/action%d", i)
}

func cloneRepo(actionRef string, dir string) error {
	strs := strings.Split(actionRef, "@")
	if len(strs) != 2 {
		return fmt.Errorf("Expected org/repo@tag, got %s", actionRef)
	}
	repoName := strs[0]
	tag := strs[1]
	cmd := exec.Command("git", "clone", "https://github.com/"+repoName, dir, "--branch", tag, "--depth", "1")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return errors.Wrap(err, "error cloning the action repo: ")
	}
	return nil
}

func runAction(dir string, stepInputs StepInputs) error {
	return nil
}

func Run(actionRef string, stepInputs StepInputs) error {
	fmt.Println("Running", actionRef, stepInputs)
	dir := tempDir()
	defer func() {
		fmt.Println("Cleaning up", dir)
		os.RemoveAll(dir)
	}()
	if err := cloneRepo(actionRef, dir); err != nil {
		return err
	}
	if err := runAction(dir, stepInputs); err != nil {
		return err
	}
	return nil
}
