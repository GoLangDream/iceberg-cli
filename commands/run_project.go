package commands

import (
	"github.com/urfave/cli/v2"
	"os"
	"os/exec"
)

func RunProject(c *cli.Context) error {
	cmd := exec.Command("air")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	return nil
}
