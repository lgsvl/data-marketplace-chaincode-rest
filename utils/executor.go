package utils

import (
	"bytes"
	"context"
	"log"
	"os"
	"os/exec"
)

//Executor executes cli commands
//go:generate counterfeiter -o ../fakes/fake_executor.go . Executor
type Executor interface { // basic host dependent functions
	Execute(command string, args []string) ([]byte, error)
}

type executor struct {
	ctx    context.Context
	logger *log.Logger
}

//NewExecutor creates a new executor
func NewExecutor(c context.Context, l *log.Logger) Executor {
	return &executor{
		ctx:    c,
		logger: l,
	}
}

func (e *executor) Execute(command string, args []string) ([]byte, error) {
	cmd := exec.Command(command, args...)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Env = os.Environ()
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	stdErr := stderr.Bytes()
	stdOut := stdout.Bytes()
	e.logger.Printf("Command-%s-executed-with-args-%#v-and-error-%s-and-output-%s", command, args, string(stdErr[:]), string(stdOut[:]))

	return stdErr, err
}
