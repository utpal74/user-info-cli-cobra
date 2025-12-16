package repository

import "errors"

// ErrNotFound returns when there is no data found in repository.
var ErrNotFound = errors.New("not found")
