package lib

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func IsInitialised(db IDb) bool {
	return db.HasHistoryStore() && db.HasMigrationStore()
}

func GetMigration(uuid string, description string, migrationType MigrationType, script IDbScript) SystemMigration {
	fileName := GetMigrationFileName(uuid, migrationType, script.GetExtension())
	return SystemMigration{
		Uuid:        uuid,
		Type:        migrationType,
		Script:      fileName,
		Description: description,
		CreatedAt:   time.Now(),
	}
}

func GetMigrationFileName(uuid string, migrationType MigrationType, extension string) string {
	return fmt.Sprintf("%s.%s.%s", uuid, migrationType.String(), extension)
}

func WriteMigrationFile(m SystemMigration) error {
	migrationFileHeader := "/*%s\n\tMIGRATION %s\n\tId: %s\n\tDescription: %s\n%s*/"
	headerDelimiter := strings.Repeat("*", 50)
	data := fmt.Sprintf(migrationFileHeader, headerDelimiter, m.GetType().String(), m.GetUuid(), m.GetDescription(), headerDelimiter)
	path := filepath.Join(Setting.GetMigrationPath(), m.GetScript())
	return ioutil.WriteFile(path, []byte(data), 0644)
}

func RemoveMigrationFile(m SystemMigration) error {
	return os.Remove(filepath.Join(Setting.GetMigrationPath(), m.GetScript()))
}

func ControlMigrationCreation(err error, tx IDbTransaction, down SystemMigration, up SystemMigration) {
	if err != nil {
		tx.Rollback()
		RemoveMigrationFile(down)
		RemoveMigrationFile(up)
	} else {
		tx.Commit()
	}
}

func ExecuteMigrations(db IDb, migrations []IMigration) error {
	var err, errTx error
	var script []byte
	var scriptPath string

	for _, v := range migrations {
		scriptPath = filepath.Join(Setting.GetMigrationPath(), v.GetScript())
		script, err = ioutil.ReadFile(scriptPath)
		if err != nil {
			break
		}

		tx := db.GetTransaction()

		executionTime, err := db.ExecuteMigration(tx, script)
		if err != nil {
			err = NewError("migration %s got an error executing script : '%s'", v.GetScript(), err.Error())
			break
		}

		err = db.UpdateMigrationHistory(v.GetId(), *executionTime)
		if err != nil {
			err = NewError("error trying to update history after execution of migration '%s'", v.GetDescription(), err.Error())
			errTx = tx.Rollback()
			break
		}

		errTx = tx.Commit()
		if errTx != nil {
			break
		}

		Success("migration \"%s\" has been executed in %f seconds", v.GetDescription(), executionTime.Seconds())
	}

	if errTx != nil {
		err = NewError("transaction error : '%s'", err.Error())
	}

	return err
}

func ReadMigrationsFromLockFile() ([]SystemMigration, error) {
	filePath := Setting.GetLockFilePath()
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var migrations []SystemMigration
	err = yaml.Unmarshal(data, &migrations)
	if err != nil {
		return nil, err
	}

	return migrations, nil
}

func WriteMigrationLockFile(down SystemMigration, up SystemMigration) error {
	filePath := Setting.GetLockFilePath()
	currentData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	var currentMigrations []SystemMigration
	err = yaml.Unmarshal(currentData, &currentMigrations)
	if err != nil {
		return err
	}

	currentMigrations = append(currentMigrations, down, up)
	newData, err := yaml.Marshal(currentMigrations)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filePath, newData, 0644)
}
