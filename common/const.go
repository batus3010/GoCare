package common

import "errors"

var (
	ErrFirstNameIsBlank = errors.New("first name cannot be blank")
	ErrLastNameIsBlank  = errors.New("last name cannot be blank")
)
