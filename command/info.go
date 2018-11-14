package command

import (
	"fmt"
	"github.com/limoli/dbshift/lib"
)

// Info is the command to get information about system
func Info() {
	fmt.Printf("%-15s\t\t%s\n", "DESCRIPTION", "PATH")
	fmt.Printf("%-15s\t\t%s\n", "ConfigPath", lib.Setting.GetConfigPath())
	fmt.Printf("%-15s\t\t%s\n", "LockFilePath", lib.Setting.GetLockFilePath())
	fmt.Printf("%-15s\t\t%s\n", "MigrationsPath", lib.Setting.GetMigrationPath())
	fmt.Printf("%-15s\t\t%s\n", "DatabaseType", lib.Setting.GetDatabaseType())
}
