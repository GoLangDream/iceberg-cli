package generate

import (
	"fmt"
	"github.com/GoLangDream/iceberg-cli/helper"
	"github.com/GoLangDream/iceberg-cli/templates"
	"github.com/GoLangDream/rgo/rstring"
	"github.com/flosch/pongo2/v5"
	"github.com/urfave/cli/v2"
	"path/filepath"
)

func Controller(c *cli.Context) error {
	name := c.Args().First()
	createController(name, c.Args().Tail()...)
	return nil
}

func createController(name string, actions ...string) {
	controllerName := rstring.Underscore(name)
	className := rstring.Camelize(name)
	helper.CreateFileFromTmpl(
		templates.ControllerTmpl(),
		fmt.Sprintf("web/controllers/%s_controller.go", controllerName),
		pongo2.Context{
			"controller_name": controllerName,
			"class_name":      className,
			"actions":         actions,
		},
	)

	for _, action := range actions {
		createView(controllerName, action)
	}
}

func createView(controller, action string) {
	actionName := rstring.Underscore(action)
	helper.MkDir(filepath.Join("web/views", controller))
	helper.CreateFileFromTmpl(
		templates.ViewTmpl(),
		fmt.Sprintf("web/views/%s/%s.html", controller, actionName),
		pongo2.Context{
			"controller_name": controller,
			"action":          action,
		},
	)
}
