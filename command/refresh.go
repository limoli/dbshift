package command

import (
	"github.com/limoli/dbshift/lib"
	"github.com/limoli/dbshift/util"
)

// Refresh is the command to discover and import new migrations
func Refresh(db lib.IDb) {
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
		util.Exit(lib.NewError(err.Error()))
	} else {
		lib.Success("%d migrations have been imported", importedMigrations/2)
	}
}
