package project

import "github.com/GoLangDream/iceberg-cli/iceberg/helper"

func CreateFile(filePath string, content ...string) {
	helper.CreateFile(Path(filePath), content...)
}
