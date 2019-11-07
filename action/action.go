package action

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type StepInputs map[string]string

func clone(actionRef string) {
	strs := strings.Split(actionRef, "@")
	if len(strs) != 2 {
		fmt.Println("Expected org/repo@tag, got", actionRef)
		os.Exit(1)
	}
	repoName := strs[0]
	tag := strs[1]
	exec.Command("git", "clone", "https://github.com/"+repoName, "--branch", tag, "--depth", "1").Run()
}

func Run(actionRef string, stepInputs StepInputs) {
	fmt.Println("Running", actionRef, stepInputs)
	clone(actionRef)
}
