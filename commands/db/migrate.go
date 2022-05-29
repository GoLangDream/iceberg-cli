package db

import (
	"github.com/GoLangDream/iceberg-cli/helper"
	"github.com/urfave/cli/v2"
)

func Migrate(c *cli.Context) error {
	err := golangMigrate().Up()
	if err != nil {
		helper.ErrorPuts("数据库迁移失败", " %s", err.Error())
	}
	return nil
}
