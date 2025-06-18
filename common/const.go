package common

import "errors"

var (
	ErrFirstNameIsBlank = errors.New("first name cannot be blank")
	ErrLastNameIsBlank  = errors.New("last name cannot be blank")
	ErrDataNotFound     = errors.New("data not found")
	ErrAddressIsBlank   = errors.New("address cannot be blank")
	ErrDataBeenDeleted  = errors.New("data has been deleted")
)
