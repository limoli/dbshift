package lib

import (
	"fmt"
	"testing"
)

func TestSuccess(t *testing.T) {
	Success("migration has been executed")
	Success("migration %s has been executed in %f seconds", "test-success", 0.24)
	Success("migration %s has been executed in %f seconds")
	Success("migration has been executed in seconds", "test-success", 0.24)
}

func TestNewError(t *testing.T) {
	const message = "this is an error"
	err := NewError("%s : %s", message)

	_, ok := err.(*libError)
	if !ok {
		t.Error("libError is not an error")
	}
}

func TestLibError_Error(t *testing.T) {
	const message = "this is an error"
	const description = "details about error"

	err := NewError("%s : %s", message, description)
	expectedValue := fmt.Sprintf("%c %s : %s", failCharacter, message, description)

	if err.Error() != expectedValue {
		t.Error("wrong error string")
	}
}
