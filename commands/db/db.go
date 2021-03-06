package db

import (
	"fmt"
	"github.com/GoLangDream/iceberg-cli/iceberg/environment"
	"github.com/GoLangDream/iceberg-cli/iceberg/helper"
	"github.com/GoLangDream/iceberg-cli/iceberg/project"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gookit/config/v2"
	"os"
)

var _golangMigrate *migrate.Migrate = nil

func golangMigrate() *migrate.Migrate {
	if _golangMigrate != nil {
		return _golangMigrate
	}

	databaseUrl := fmt.Sprintf(
		"mysql://%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		config.String(keyWithEnv("database.%s.username")),
		config.String(keyWithEnv("database.%s.password")),
		config.String(keyWithEnv("database.%s.host")),
		config.String(keyWithEnv("database.%s.port")),
		config.String(keyWithEnv("database.%s.name")),
		config.String(keyWithEnv("database.%s.encoding")),
	)

	m, err := migrate.New(
		fmt.Sprintf("file://%s", project.Path(project.MigrationPath())),
		databaseUrl,
	)
	if err != nil {
		helper.ErrorPuts("连接数据库失败", "请确认数据库配置正确 %s", err)
		os.Exit(0)
	}

	m.Log = &helper.MigrationLogger{}

	_golangMigrate = m
	return _golangMigrate
}

func keyWithEnv(format string) string {
	return fmt.Sprintf(format, environment.Name())
}
