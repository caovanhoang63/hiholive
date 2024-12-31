package authmodel

import "errors"

var (
	ErrPinAlreadyExists = errors.New("pin already exists")
	ErrInvalidPin       = errors.New("invalid pin")
)
