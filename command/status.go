package command

import (
	"fmt"
	"github.com/limoli/dbshift/lib"
)

// Status is the command to get status about migrations
func Status(db lib.IDb) {
	counterMigrationUpgrade, counterMigrationDowngrade := db.GetStatus()
	fmt.Printf("↑ %d migrations to upgrade\n", counterMigrationUpgrade)
	fmt.Printf("↓ %d migrations to downgrade\n", counterMigrationDowngrade)
}
