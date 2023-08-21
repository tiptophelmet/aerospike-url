package aerourl

import "errors"

var (
	ErrEmptyConnStr  = errors.New("empty aerospike connection string")
	ErrNilURL        = errors.New("aerospike URL must be initialized with connection string")
	ErrInvalidScheme = errors.New("invalid url scheme, want: aerospike://")
	ErrEmptyHostname = errors.New("aerospike hostname cannot be empty")
	ErrEmptyPort     = errors.New("aerospike port cannot be empty")
)
