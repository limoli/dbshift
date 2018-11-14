package command

import (
	"github.com/limoli/dbshift/lib"
	"github.com/limoli/dbshift/util"
)

// Execute is the command to execute migrations upgrades or downgrades
func Execute(db lib.IDb, migrationType lib.MigrationType) {
	migrations, err := db.GetMigrationsByType(migrationType, nil)
	if err != nil {
		util.Exit(err)
	}

	err = lib.ExecuteMigrations(db, migrations)
	if err != nil {
		util.Exit(err)
	}

	lib.Success("all migrations have been executed successfully")
}
