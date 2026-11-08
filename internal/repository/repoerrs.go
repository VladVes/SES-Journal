package repository

import "errors"

var (
	ErrRecordNotFoundByID = errors.New("record not found by given ID")
)
