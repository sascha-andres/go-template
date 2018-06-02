package wrapper

import (
	"os"
	"os/exec"
	"syscall"

	"github.com/Sirupsen/logrus"
	"github.com/pkg/errors"
)

var (
	gitExecutable string
)

func init() {
	var err error
	gitExecutable, err = exec.LookPath("git")
	if err != nil {
		panic(err)
	}
}

func Git(args ...string) (int, error) {
	logrus.
		WithField("package", "wrapper").
		WithField("method", "Git").
		Debug("git ", args)
	command := exec.Command(gitExecutable, args...)
	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr
	return startAndWait(command)
}

// startAndWait calls the command and returns the result
func startAndWait(command *exec.Cmd) (int, error) {
	var err error
	if err = command.Start(); err != nil {
		return -1, errors.Wrap(err, "could not start command")
	}
	err = command.Wait()
	if err == nil {
		return 0, nil
	}
	if exitError, ok := err.(*exec.ExitError); ok {
		if err.(*exec.ExitError).Stderr == nil {
			return 0, nil
		}
		if status, ok := exitError.Sys().(syscall.WaitStatus); ok {
			return status.ExitStatus(), errors.Wrap(err, "error waiting for command")
		}
	} else {
		return -1, errors.Wrap(err, "error waiting for command")
	}
	return 0, nil
}
