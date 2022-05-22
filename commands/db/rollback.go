package db

import (
	"github.com/GoLangDream/iceberg-cli/iceberg/helper"
	"github.com/urfave/cli/v2"
)

func Rollback(c *cli.Context) error {
	err := golangMigrate().Steps(-1)
	if err != nil {
		helper.ErrorPuts("数据库回滚失败 %s", err.Error())
	}
	return err
}
