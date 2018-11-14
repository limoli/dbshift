package command

import (
	"fmt"
	"github.com/limoli/dbshift/lib"
	"github.com/limoli/dbshift/util"
)

// Migrations is the command to get migrations to upgrade or downgrade
func Migrations(db lib.IDb, migrationType lib.MigrationType) {
	migrations, err := db.GetMigrationsByType(migrationType, nil)
	if err != nil {
		util.Exit(err)
	}

	fmt.Printf("UUID%32s\t\tDESCRIPTION\n", "")
	for _, v := range migrations {
		fmt.Printf("%36s\t\t%s\n", v.GetUuid(), v.GetDescription())
	}
}
