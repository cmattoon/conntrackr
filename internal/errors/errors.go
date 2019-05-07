package errors

import (
	"log"
)

type appError struct {
	text string
	code int
}

func (a *appError) Error() string {
	return a.text
}

func (a *appError) Code() int {
	return a.code
}

func AppError(s string, i int) error {
	return &appError{
		text: s,
		code: i,
	}
}

func Check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
