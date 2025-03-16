package internal

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
)

type Command struct {
	args []string
}

func RunCommand(args ...string) *Command {
	return &Command{
		args,
	}
}

func (c *Command) runCommand() (string, int, error) {
	cmd := exec.Command(c.args[0], c.args[1:]...)
	out, err := cmd.CombinedOutput()
	var cleanOut = strings.TrimSuffix(string(out), "\n")
	if err != nil {
		return cleanOut, cmd.ProcessState.ExitCode(), fmt.Errorf("error executing %s: %v", c.args[0], err)
	}
	return cleanOut, cmd.ProcessState.ExitCode(), nil
}

func (c *Command) Single() (string, int, error) {
	return c.runCommand()
}

func (c *Command) Multi(wg *sync.WaitGroup) (string, int, error) {
	defer wg.Done()
	return c.runCommand()
}

func ExitError(message any, code int) {
	fmt.Fprintf(os.Stderr, "pkyu: \033[31m%s\033[0m\n", message)
	os.Exit(code)
}

func BinaryExists(name string) (bool, error) {
	for range strings.SplitSeq(os.Getenv("PATH"), ":") {
		path, err := exec.LookPath(name)
		if err == nil && path != "" {
			return true, nil
		}
		return false, fmt.Errorf("could not search $PATH for %s: %v", name, err)
	}
	return false, nil
}
