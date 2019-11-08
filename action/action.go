package action

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"strings"
)

type StepInputs map[string]string

func tempDir() string {
	i, err := rand.Int(rand.Reader, big.NewInt(100000000))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return fmt.Sprintf("/tmp/action%d", i)
}

func cloneRepo(actionRef string, dir string) {
	strs := strings.Split(actionRef, "@")
	if len(strs) != 2 {
		fmt.Println("Expected org/repo@tag, got", actionRef)
		os.Exit(1)
	}
	repoName := strs[0]
	tag := strs[1]
	cmd := exec.Command("git", "clone", "https://github.com/"+repoName, dir, "--branch", tag, "--depth", "1")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Run(actionRef string, stepInputs StepInputs) {
	fmt.Println("Running", actionRef, stepInputs)
	dir := tempDir()
	defer func() {
		fmt.Println("Cleaning up", dir)
		os.RemoveAll(dir)
	}()
	cloneRepo(actionRef, dir)
}
