package errors

import "errors"

var (
	ErrCategoryNotFound    = errors.New("category not found")
	ErrDuplicateCategory   = errors.New("category already exists")
	ErrTransactionNotFound = errors.New("transaction not found")
	ErrValidation          = errors.New("validation failed")
)
