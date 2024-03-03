package cmd

import (
	"os"
	"socialForumBackend/cmd/app"
	config "socialForumBackend/internal/config"
	db "socialForumBackend/internal/database"
	"socialForumBackend/internal/server"

	"github.com/urfave/cli"
)

func main() {
	app.Init()
	defer app.ShutDown()

	cliApp := cli.NewApp()
	cliApp.Name = "forumApp"
	cliApp.Version = "1.0.0"
	cliApp.Usage = ""

	migrationConfig := db.MigrationConfig{
		Driver: "postgres",
		URL:    config.Database.ConnectionURL(),
	}

	cliApp.Commands = cli.Commands{
		{
			Name:  "server",
			Usage: "Start server",
			Action: func(c *cli.Context) error {
				return server.Start()
			},
		},
		{
			Name:  "migrate",
			Usage: "run db migrations",
			Action: func(c *cli.Context) error {
				return db.RunDatabaseMigrations(&migrationConfig)
			},
		},
		{
			Name:  "rollback",
			Usage: "rollback db migrations",
			Action: func(c *cli.Context) error {
				return db.RollbackLatestMigration(&migrationConfig)
			},
		},
	}

	if err := cliApp.Run(os.Args); err != nil {
		panic(err)
	}

}
