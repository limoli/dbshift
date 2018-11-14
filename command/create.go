package command

import (
	"github.com/gofrs/uuid"
	"github.com/limoli/dbshift/lib"
	"github.com/limoli/dbshift/util"
)

// Create is the command to create migrations
func Create(db lib.IDb, description string) {
	mUuid, err := uuid.NewV4()
	if err != nil {
		util.Exit(err)
	}

	// Generate migration objects (downgrade and upgrade)
	uuidString := mUuid.String()
	migrationDowngrade := lib.GetMigration(uuidString, description, lib.Downgrade, db)
	migrationUpgrade := lib.GetMigration(uuidString, description, lib.Upgrade, db)

	// Write migration file (downgrade)
	err = lib.WriteMigrationFile(migrationDowngrade)
	if err != nil {
		util.Exit(err)
	}

	// Write migration file (upgrade)
	err = lib.WriteMigrationFile(migrationUpgrade)
	if err != nil {
		lib.RemoveMigrationFile(migrationDowngrade)
		util.Exit(err)
	}

	// Create migration into db with transaction
	tx := db.GetTransaction()

	defer lib.ControlMigrationCreation(err, tx, migrationDowngrade, migrationUpgrade)

	err = db.CreateMigration(tx, &migrationDowngrade)
	if err != nil {
		util.Exit(err)
	}

	err = db.CreateMigration(tx, &migrationUpgrade)
	if err != nil {
		util.Exit(err)
	}

	// Create lock file with migration (downgrade and upgrade information)
	err = lib.WriteMigrationLockFile(migrationDowngrade, migrationUpgrade)
	if err != nil {
		util.Exit(err)
	}

	lib.Success("migrations have been created")
}
