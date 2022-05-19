package core

import (
	"LyGen/commands"
	"LyGen/constant"
	"github.com/urfave/cli/v2"
)

func usage(app *cli.App) {
	app.Name = constant.ProjectName // 项目名称
	app.Authors = []*cli.Author{{
		Name:  constant.Author,
		Email: constant.AuthorEmail,
	}}
	app.Version = constant.Version                                                // 版本号
	app.Copyright = "@Copyright " + constant.Copyright                            // 版权保护
	app.Usage = "Quickly generate CURD and documentation for operating MYSQL.etc" // 说明
	app.Commands = []*cli.Command{
		{
			Name:    "help",
			Aliases: []string{"h", "?"},
			Usage:   "show help",
			Action: func(c *cli.Context) error {
				_ = cli.ShowAppHelp(c)
				return nil
			},
		},
		{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "print the version",
			Action: func(c *cli.Context) error {
				cli.ShowVersion(c)
				return nil
			},
		},
	}
	app.HideVersion = true
	app.HideHelp = true
	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: "mysql", Value: "", Usage: "Database address"},
	}
}


func Start(app *cli.App) {
	usage(app)
	app.Action = func(c *cli.Context) error {
		constant.MysqlDataSource = c.String("mysql")
		cmd := new(commands.Commands)
		if err := cmd.Start(); err != nil {
			return cli.Exit(err.Error(), 1)
		}
		return nil
	}
}
