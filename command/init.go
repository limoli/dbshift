package command

import (
	"fmt"
	"github.com/limoli/dbshift/lib"
	"github.com/limoli/dbshift/util"
)

// Init is the command to initialise the system
func Init(db lib.IDb) {
	if !db.HasMigrationStore() {
		if err := db.CreateMigrationStore(); err != nil {
			util.Exit(err)
		}
	}

	if !db.HasHistoryStore() {
		if err := db.CreateHistoryStore(); err != nil {
			util.Exit(err)
		}
	}

	lib.Success("database has been initialised")

	migrations, err := lib.ReadMigrationsFromLockFile()
	if err != nil {
		util.Exit(err)
	}

	lib.Success("%d migrations have been detected in lock file", len(migrations)/2)

	iMigrations := make([]lib.IMigration, 0)
	for k := range migrations {
		iMigrations = append(iMigrations, &migrations[k])
	}

	importedMigrations, err := db.ImportMigrations(iMigrations)
	if err != nil {
		fmt.Println(err)
	} else {
		lib.Success("%d migrations have been imported", importedMigrations/2)
	}
}
