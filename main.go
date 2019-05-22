package main

import (
	"errors"
	"fmt"
	"github.com/abiosoft/ishell"
	"github.com/limoli/dbshift/command"
	"github.com/limoli/dbshift/lib"
	"github.com/limoli/dbshift/lib/mysql"
	"github.com/limoli/dbshift/util"
	"os"
)

func main() {
	// Get database instance according to database type
	db, err := getDatabase(lib.Setting.GetDatabaseType())
	if err != nil {
		util.Exit(err)
	}

	// Try connection to database
	err = db.Connect(lib.Setting.GetDatabaseConnection())
	if err != nil {
		util.Exit(lib.NewError("impossible to connect to db : '%s'", err.Error()))
	}

	// Run shell
	shell := ishell.New()

	commands := getShellCommands(db)
	for k := range commands {
		shell.AddCmd(commands[k])
	}

	if len(os.Args) > 1 {
		err = shell.Process(os.Args[1:]...)
		if err != nil {
			util.Exit(err)
		}
	} else {
		shell.Run()
	}
}

func getDatabase(dbType lib.DatabaseType) (lib.IDb, error) {
	var db lib.IDb

	switch dbType {
	case lib.DatabaseTypeMySQL:
		db = new(mysql.Db)
		break
	default:
		errMsg := fmt.Sprintf("db type \"%s\" not available %v\n", dbType, lib.AvailableDatabases)
		return nil, errors.New(errMsg)
	}

	return db, nil
}

func getShellCommands(db lib.IDb) []*ishell.Cmd {
	return []*ishell.Cmd{
		{
			Name:     "init",
			Help:     "init",
			LongHelp: "Install dbshift using your database.",
			Func: func(c *ishell.Context) {
				util.ExitIfInitialised(db)
				command.Init(db)
			},
		}, {
			Name:     "refresh",
			Help:     "refresh",
			LongHelp: "It discovers and imports new migrations.",
			Func: func(c *ishell.Context) {
				util.ExitIfNotInitialised(db)
				command.Refresh(db)
			},
		}, {
			Name:     "info",
			Help:     "info",
			LongHelp: "It returns information about dbshift.",
			Func: func(c *ishell.Context) {
				util.ExitIfNotInitialised(db)
				command.Info()
			},
		}, {
			Name:     "status",
			Help:     "status",
			LongHelp: "It returns the current status of database along migrations.",
			Func: func(c *ishell.Context) {
				util.ExitIfNotInitialised(db)
				command.Status(db)
			},
		}, {
			Name:     "create",
			Help:     "create <description>",
			LongHelp: "It creates a migration with description.",
			Func: func(c *ishell.Context) {
				util.ExitIfNotInitialised(db)
				if len(c.Args) != 1 {
					util.Exit(lib.NewError("missing migration description"))
				}
				description := c.Args[0]
				command.Create(db, description)
			},
		}, {
			Name:     "migrations-upgrade",
			Help:     "migrations-upgrade",
			LongHelp: "It returns the list of migrations eligible to upgrade.",
			Func: func(c *ishell.Context) {
				util.ExitIfNotInitialised(db)
				command.Migrations(db, lib.Upgrade)
			},
		}, {
			Name:     "migrations-downgrade",
			Help:     "migrations-downgrade",
			LongHelp: "It returns the list of migrations eligible to downgrade.",
			Func: func(c *ishell.Context) {
				util.ExitIfNotInitialised(db)
				command.Migrations(db, lib.Downgrade)
			},
		}, {
			Name:     "upgrade",
			Help:     "upgrade",
			LongHelp: "It upgrades all the migrations eligible to upgrade.",
			Func: func(c *ishell.Context) {
				util.ExitIfNotInitialised(db)
				command.Execute(db, lib.Upgrade)
			},
		}, {
			Name:     "downgrade",
			Help:     "downgrade",
			LongHelp: "It downgrades all the migrations eligible to downgrade.",
			Func: func(c *ishell.Context) {
				util.ExitIfNotInitialised(db)
				command.Execute(db, lib.Downgrade)
			},
		},
	}
}
