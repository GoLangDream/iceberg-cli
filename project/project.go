package project

import "github.com/GoLangDream/iceberg-cli/iceberg/environment"

func Init() {
	initConfig()
	initPongo2()
	environment.Init()
}
