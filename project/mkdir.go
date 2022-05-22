package project

import "github.com/GoLangDream/iceberg-cli/iceberg/helper"

func MkDir(name string) {
	helper.MkDir(Path(name))
}
