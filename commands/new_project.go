package commands

import (
	_ "embed"
	"github.com/GoLangDream/iceberg-cli/helper"
	"github.com/GoLangDream/iceberg-cli/templates"
	"github.com/GoLangDream/rgo/rstring"
	"github.com/flosch/pongo2/v5"
	"github.com/urfave/cli/v2"
	"os"
)

var projectName string
var className string

func NewProject(c *cli.Context) error {
	name := c.Args().First()
	projectName = rstring.Underscore(name)
	className = rstring.Camelize(name)
	createProject()
	return nil
}

func projectInfoContext() pongo2.Context {
	return pongo2.Context{
		"project_name": projectName,
		"class_name":   className,
	}
}

func createProject() {
	helper.MkDir(projectName)
	err := os.Chdir(projectName)
	if err != nil {
		panic(err)
	}
	templates.CreateDirs(createFile, createDir)
	initGoMod()
}

func initGoMod() {
	helper.RunCmd("go", "mod", "init", projectName)
	helper.RunCmd("go", "get", "github.com/GoLangDream/iceberg")
	helper.RunCmd("go", "mod", "tidy")

	helper.RunCmd("go", "install", "github.com/cosmtrek/air@latest")
	helper.RunCmd("air", "init")
}

func createDir(path string) {
	helper.MkDir(path)
}

func createFile(path string) {
	helper.CreateFileFromTmpl(
		templates.File(path+".tmpl"),
		path,
		projectInfoContext(),
	)
}
