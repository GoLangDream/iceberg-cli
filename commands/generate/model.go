package generate

import (
	"fmt"
	"github.com/GoLangDream/iceberg-cli/iceberg/helper"
	"github.com/GoLangDream/iceberg-cli/iceberg/templates"
	"github.com/GoLangDream/rgo/rstring"
	"github.com/flosch/pongo2/v5"
	"github.com/urfave/cli/v2"
	"strings"
)

func Model(c *cli.Context) error {
	name := c.Args().First()
	createModel(name, c.Args().Tail()...)
	return nil
}

type ColumnInfo struct {
	Name  string
	CType string
}

func createModel(name string, columns ...string) {
	modelName := rstring.Underscore(name)
	className := rstring.Camelize(name)

	columnInfos := []ColumnInfo{}

	for _, column := range columns {
		columnInfos = append(columnInfos, getColumnInfo(column))
	}

	helper.CreateFileFromTmpl(
		templates.ModelTmpl(),
		fmt.Sprintf("models/%s.go", modelName),
		pongo2.Context{
			"class_name": className,
			"columns":    columnInfos,
		},
	)
}

func getColumnInfo(str string) ColumnInfo {
	info := strings.Split(str, ":")
	return ColumnInfo{rstring.Camelize(info[0]), info[1]}
}
