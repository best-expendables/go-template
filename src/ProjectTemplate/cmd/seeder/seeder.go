package main

import (
	"fmt"

	"${REPO_HOST}/${PROJ_NAME}/cmd/seeder/seeders"
	"${REPO_HOST}/${PROJ_NAME}/config"
	"${REPO_HOST}/${PROJ_NAME}/connection"
)

func main() {
	appConf := config.GetAppConfigFromEnv()
	postgresConnection := connection.NewPostgresConnection(appConf.DBConfig)
	db := postgresConnection.CreateDB()
	runner([]Seeder{
		seeders.New${SERVICE_NAME}(db),
	})
}

type Seeder interface {
	ShouldRun() bool
	Prepare() error
	Run() error
	Commit() error
	Rollback() error
}

func runner(list []Seeder) {
	for _, s := range list {
		if !s.ShouldRun() {
			continue
		}
		fmt.Printf("seeding %T\n", s)
		if err := s.Prepare(); err != nil {
			panic(err)
		}
		if err := s.Run(); err != nil {
			if err := s.Rollback(); err != nil {
				fmt.Println("Rollback failed", err)
			}
			panic(err)
		}
		if err := s.Commit(); err != nil {
			fmt.Println("Commit failed", err)
		}
	}
}
