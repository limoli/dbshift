package lib

import "fmt"

const successCharacter rune = '✔'
const failCharacter rune = '✘'

type libError struct {
	msg string
}

func (e *libError) Error() string {
	return fmt.Sprintf("%c %s", failCharacter, e.msg)
}

func NewError(text string, args ...interface{}) error {
	if len(args) > 0 {
		text = fmt.Sprintf(text, args...)
	}
	return &libError{
		msg: text,
	}
}

func Success(text string, args ...interface{}) {
	if len(args) > 0 {
		text = fmt.Sprintf(text, args...)
	}
	fmt.Printf("%c %s\n", successCharacter, text)
}
