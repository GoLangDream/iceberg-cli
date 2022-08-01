package project

import "github.com/GoLangDream/iceberg-cli/iceberg/helper"

func MkDir(name string) {
	helper.MkDir(Path(name))
}

func MkDirIfNotExists(name string) {
	helper.MkDirIfNotExists(Path(name))
}
