package constants

import "errors"

var (
	ErrNotFound             = errors.New("record not found")
	ErrBalanceNotSufficient = errors.New("balance is not sufficient")
)
