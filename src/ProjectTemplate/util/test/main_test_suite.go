package test

import (
	"${REPO_HOST}/${PROJ_NAME}/cmd/migrate/migrations"
	"${REPO_HOST}/${PROJ_NAME}/config"
	"${REPO_HOST}/${PROJ_NAME}/connection"
	"errors"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/pressly/goose"
	"github.com/stretchr/testify/suite"
	"github.com/subosito/gotenv"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
	"time"
)

type MainTestSuite struct {
	suite.Suite
	DB             *gorm.DB
	appConf        config.AppConfig
	dockerResource *DockerResource
}

type DockerResource struct {
	name     string
	pool     *dockertest.Pool
	resource *dockertest.Resource
}

var (
	envFile = "../../docker/Application/env_files/main.env"
)

func (s *MainTestSuite) SetupSuite() {
	err := gotenv.OverLoad("../../docker/Application/env_files/main.env")
	if err != nil {
		log.Fatalf("Could not open enviroment file %s, %s", envFile, err)
	}
	s.appConf = config.GetAppConfigFromEnv()
	s.setupDocker()
	s.setupMigration()
}

func (s *MainTestSuite) TearDownSuite() {
	// After done, kill and remove the container
	if err := s.dockerResource.pool.RemoveContainerByName(s.dockerResource.name); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
}

func (s *MainTestSuite) setupDocker() {
	pool, err := dockertest.NewPool("")
	dbConf := s.appConf.DBConfig
	dbConf.DbHost = "localhost"
	dbConf.Slaves = nil

	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	s.dockerResource = &DockerResource{}
	s.dockerResource.name = dbConf.DbName + "_test_" + strconv.Itoa(time.Now().Nanosecond())
	opts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "9.6.2",
		Name:       s.dockerResource.name,
		Env: []string{
			"POSTGRES_USER=" + dbConf.DbUser,
			"POSTGRES_PASSWORD=" + dbConf.DbPass,
			"POSTGRES_DB=" + dbConf.DbName,
		},
		ExposedPorts: []string{"5432"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432": {
				{HostPort: "5432"},
			},
		},
	}

	resource, err := pool.RunWithOptions(&opts)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	if err = pool.Retry(func() (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = errors.New("")
			}
		}()
		postgresConnection := connection.NewPostgresConnection(dbConf)
		db := postgresConnection.CreateDB()
		s.DB = db
		rawDB, err := db.DB()
		if err != nil {
			return err
		}
		return rawDB.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	s.dockerResource.pool = pool
	s.dockerResource.resource = resource
}

func (s *MainTestSuite) setupMigration() {
	migrations.SetDB(s.DB)
	migrationPath, err := s.GetMigrationPath()
	rawDB, err := s.DB.DB()
	if err != nil {
		log.Fatalf("Could not apply migration", err)
	}
	err = goose.Run("up", rawDB, migrationPath)
	if err != nil {
		log.Fatalf("Could not apply migration", err)
	}
}

// GetMigrationPath check the migration directories
func (s *MainTestSuite) GetMigrationPath() (string, error) {
	var migrationPaths = []string{
		"../cmd/migrate/migrations",
		"../../cmd/migrate/migrations",
	}

	for _, path := range migrationPaths {
		if exist, err := s.folderExists(path); exist == true {
			return path, err
		}
	}

	return "", errors.New("Cannot find migration folder")
}

// Check if folder exists
func (s *MainTestSuite) folderExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
