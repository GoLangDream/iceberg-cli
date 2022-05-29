package main

import (
	"github.com/GoLangDream/iceberg-cli/commands"
	"github.com/GoLangDream/iceberg-cli/commands/db"
	"github.com/GoLangDream/iceberg-cli/commands/generate"
	"github.com/GoLangDream/iceberg-cli/helper"
	"github.com/GoLangDream/iceberg-cli/project"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	project.Init()

	app := &cli.App{
		Name:  "iceberg",
		Usage: "iceberg 命令行工具",
		Commands: []*cli.Command{
			{
				Name:      "new",
				Usage:     "创建项目",
				ArgsUsage: "NAME",
				Action:    commands.NewProject,
			},
			{
				Name:   "run",
				Usage:  "启动程序",
				Action: commands.RunProject,
			},
			{
				Name:   "db:migrate",
				Usage:  "迁移数据库",
				Action: db.Migrate,
			},
			{
				Name:   "db:rollback",
				Usage:  "回滚数据迁移",
				Action: db.Rollback,
			},
			{
				Name:    "generate",
				Aliases: []string{"g"},
				Usage:   "生成新的代码",
				Subcommands: []*cli.Command{
					{
						Name:      "migration",
						Aliases:   []string{"m"},
						Usage:     "创建一个数据库迁移任务",
						ArgsUsage: "NAME",
						Action:    generate.Migration,
					},
					{
						Name:      "controller",
						Aliases:   []string{"c"},
						Usage:     "创建一个controller",
						ArgsUsage: "NAME ACTION...",
						Action:    generate.Controller,
					},
					{
						Name:      "model",
						Usage:     "创建一个model",
						ArgsUsage: "NAME COLUMN:TYPE",
						Action:    generate.Model,
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	helper.RunCmd("go", "fmt", "./...")
}
