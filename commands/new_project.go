package commands

import (
	"github.com/GoLangDream/iceberg-cli/iceberg/helper"
	"github.com/urfave/cli/v2"
	"os"
)

func NewProject(c *cli.Context) error {
	projectName := c.Args().First()
	createProject(projectName)
	return nil
}

func createProject(name string) {
	helper.MkDir(name)
	os.Chdir(name)
}
