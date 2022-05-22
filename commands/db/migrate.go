package db

import "github.com/urfave/cli/v2"

func Migrate(c *cli.Context) error {
	golangMigrate().Up()
	return nil
}
