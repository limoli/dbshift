package lib

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	Setting *setting
)

var (
	errBadConfigPath    = errors.New(configPathEnv + " is not set")
	errBadMigrationPath = errors.New("migration path is not valid")
)

const (
	configPathEnv  string = "DBSHIFT_CONFIG"
	configFileName string = "dbshift.yaml"
	lockFileName   string = "dbshift.lock.yaml"
)

type setting struct {
	configPath    string
	lockFilePath  string
	migrationPath string
	dbType        DatabaseType
	dbConnection  IDbConfig
}

func (s *setting) GetConfigPath() string {
	return s.configPath
}

func (s *setting) GetLockFilePath() string {
	return s.lockFilePath
}

func (s *setting) GetMigrationPath() string {
	return s.migrationPath
}

func (s *setting) GetDatabaseType() DatabaseType {
	return s.dbType
}

func (s *setting) GetDatabaseConnection() IDbConfig {
	return s.dbConnection
}

func init() {
	// Get configuration path from environment variable
	configPath := os.Getenv(configPathEnv)
	if configPath == "" {
		exit(errBadConfigPath)
	}

	// Read yaml configuration file
	config, err := readConfiguration(configPath, configFileName)
	if err != nil {
		exit(err)
	}

	// Get migration path from environment variable
	migrationPath := config.Db.Migration.Path
	if config.Db.Migration.PathEnv != "" {
		migrationPath = os.Getenv(config.Db.Migration.PathEnv)
	}

	// Check migration path if exists
	err = checkMigrationsPath(migrationPath)
	if err != nil {
		exit(err)
	}

	// Check lock file
	err = checkLockFile(configPath, lockFileName)
	if err != nil {
		exit(err)
	}

	Setting = &setting{
		configPath:    configPath,
		lockFilePath:  filepath.Join(configPath, lockFileName),
		migrationPath: migrationPath,
		dbType:        config.Db.Type,
		dbConnection:  &config.Db.Connection,
	}
}

func exit(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func readConfiguration(configPath string, configFileName string) (*Configuration, error) {
	filePath := filepath.Join(configPath, configFileName)

	configData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	config := new(Configuration)
	err = yaml.Unmarshal(configData, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func checkMigrationsPath(migrationsPath string) error {
	_, err := os.Stat(migrationsPath)
	if err != nil {
		return errBadMigrationPath
	}
	return nil
}

func checkLockFile(configPath string, lockFileName string) error {
	filePath := filepath.Join(configPath, lockFileName)
	_, err := ioutil.ReadFile(filePath)
	if err != nil {
		return ioutil.WriteFile(filePath, []byte(""), 0644)
	}
	return nil
}
