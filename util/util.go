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

// Exit causes the current program to exit with code 1 printing the error
func Exit(err error) {
	fmt.Println(err)
	os.Exit(1)
}

// ExitIfInitialised causes the current program to exit if system is already initialised
func ExitIfInitialised(db lib.IDb) {
	if lib.IsInitialised(db) {
		Exit(errAlreadyInitialised)
	}
}

// ExitIfNotInitialised causes the current program to exit if system is not initialised
func ExitIfNotInitialised(db lib.IDb) {
	if !lib.IsInitialised(db) {
		Exit(errNotInitialised)
	}
}
