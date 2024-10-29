package models

import (
	"errors"
)

var (
	ErrNoRecord = errors.New("models: no matching record")

	ErrInvalidCredentials = errors.New("models: invalid credentials")

	ErrDuplicateEmails = errors.New("models: duplicate email")
)
