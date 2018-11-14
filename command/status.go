package command

import (
	"fmt"
	"github.com/limoli/dbshift/lib"
	"github.com/limoli/dbshift/util"
)

// Status is the command to get status about migrations
func Status(db lib.IDb) {
	counterMigrationUpgrade, err := db.GetExecutableMigrationsCounter(lib.Upgrade)
	if err != nil {
		util.Exit(err)
	}

	counterMigrationDowngrade, err := db.GetExecutableMigrationsCounter(lib.Downgrade)
	if err != nil {
		util.Exit(err)
	}

	fmt.Printf("↑ %d migrations to upgrade\n", counterMigrationUpgrade)
	fmt.Printf("↓ %d migrations to downgrade\n", counterMigrationDowngrade)
}
