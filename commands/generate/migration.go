package generate

import (
	"fmt"
	"github.com/GoLangDream/iceberg-cli/project"
	"github.com/urfave/cli/v2"
	"path/filepath"
	"time"
)

func Migration(c *cli.Context) error {
	name := c.Args().First()
	createMigration(name)
	return nil
}

func version() string {
	return time.Now().Format("20060102150405")
}

func createMigration(name string) {
	version := version()
	upFileName := fmt.Sprintf("%s_%s.up.sql", version, name)
	downFileName := fmt.Sprintf("%s_%s.down.sql", version, name)
	project.CreateFile(filepath.Join(project.MigrationPath(), upFileName))
	project.CreateFile(filepath.Join(project.MigrationPath(), downFileName))
}
