package util

import (
	"fmt"
	"github.com/limoli/dbshift/lib"
	"os"
)

var (
	errAlreadyInitialised = lib.NewError("db is already initialised")
	errNotInitialised     = lib.NewError("db not initialised")
)

func Exit(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func ExitIfInitialised(db lib.IDb) {
	if lib.IsInitialised(db) {
		Exit(errAlreadyInitialised)
	}
}

func ExitIfNotInitialised(db lib.IDb) {
	if !lib.IsInitialised(db) {
		Exit(errNotInitialised)
	}
}
