package server

import (
	"github.com/bitxeno/go-docker-skeleton/internal/app"
	"github.com/bitxeno/go-docker-skeleton/web"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

var (
	flags = []cli.Flag{
		&cli.IntFlag{
			Name:    "port",
			Aliases: []string{"p"},
			Usage:   "Define an alternate web server port",
		},
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "Load configuration from `FILE`",
		},
		&cli.BoolFlag{
			Name:    "debug",
			Aliases: []string{"vv"},
			Usage:   "Enable debug output",
			Value:   false,
		},
		&cli.BoolFlag{
			Name:    "verbose",
			Aliases: []string{"vvv"},
			Usage:   "Enable verbose output",
			Value:   false,
		},
	}

	Command = &cli.Command{
		Name:   "server",
		Usage:  "Run web server",
		Flags:  flags,
		Action: action,
	}
)

func action(c *cli.Context) error {
	// init config
	debug := false
	if c.Bool("debug") || c.Bool("verbose") {
		debug = true
	}
	conf, err := app.InitConfig(c.String("config"), debug)
	if err != nil {
		return err
	}
	if err = app.InitSettings(conf, debug); err != nil {
		return err
	}

	// init logger
	if c.Bool("debug") {
		conf.Log.Level = "debug"
		conf.Db.Debug = true
	}
	if c.Bool("verbose") {
		conf.Log.Level = "trace"
		conf.Db.Debug = true
	}
	if err := app.InitLogger(conf); err != nil {
		return err
	}

	// init db
	if err := app.InitDb(conf); err != nil {
		return err
	}

	printVersion()
	port := conf.Server.Port
	if c.Int("port") > 0 {
		port = c.Int("port")
	}
	return web.Run(conf.Server.ListenAddr, port)
}

func printVersion() {
	color.New(color.FgGreen).Print("Starting server version: ")
	color.New(color.FgCyan).Printf("%s@%s@%v\n", app.Version.Version, app.Version.BuildDate, app.Mode)
}
