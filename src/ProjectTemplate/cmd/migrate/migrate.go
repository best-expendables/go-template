package main

import (
	"${REPO_HOST}/${PROJ_NAME}/cmd/migrate/migrations"
	"${REPO_HOST}/${PROJ_NAME}/config"
	"${REPO_HOST}/${PROJ_NAME}/connection"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"log"
	"os"
)

var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
	dir   = flags.String("dir", "cmd/migrate/migrations", "directory with migration files")
)

func main() {
	flags.Usage = usage
	flags.Parse(os.Args[1:])

	args := flags.Args()
	if len(args) < 1 {
		flags.Usage()
		return
	}
	command := args[0]
	if command == "-h" || command == "--help" {
		flags.Usage()
		return
	}
	appConf := config.GetAppConfigFromEnv()
	postgresConnection := connection.NewPostgresConnection(appConf.DBConfig)
	db := postgresConnection.CreateDB()
	rawDb := postgresConnection.GetRawDB(db)

	migrations.SetDB(db)
	if err := goose.Run(command, rawDb, *dir, args[1:]...); err != nil {
		log.Fatalf("goose run: %v", err)
	}
}

func usage() {
	fmt.Print(usagePrefix)
	flags.PrintDefaults()
	fmt.Print(usageCommands)
	fmt.Print(usageCreate)
}

var (
	usagePrefix = `Usage: migrate [OPTIONS] COMMAND

Options:
`

	usageCommands = `
Commands:
    up         Migrate the DB to the most recent version available
    up-by-one  Migrate the DB to the next available version
    down       Roll back the version by 1
    refresh    Down all migrations and apply them again
    status     Dump the migration status for the current DB
    version    Print the current version of the database
    create     Creates a blank migration template
`

	usageCreate = `create must be of form: migrate [OPTIONS] create NAME [go|sql|gorm]`
)
