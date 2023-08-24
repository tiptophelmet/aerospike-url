// Aerospike URL package.
// Aerospike connection string is validated & processed into [aerourl.AerospikeURL] here.
package aerourl

import "errors"

var (
	// Empty connection string passed to [aerourl.Init]
	ErrEmptyConnStr = errors.New("empty aerospike connection string")

	// Aerospike URL [aerourl.AerospikeURL] with nil internal [net/url.URL]
	ErrNilURL = errors.New("aerospike URL must be initialized with connection string")

	// Aerospike URL validation failed due to invalid URL scheme
	ErrInvalidScheme = errors.New("invalid url scheme, want: aerospike://")

	// Aerospike URL validation failed due to empty Aerospike DB hostname
	ErrEmptyHostname = errors.New("aerospike hostname cannot be empty")

	// Aerospike URL validation failed due to empty Aerospike DB port
	ErrEmptyPort = errors.New("aerospike port cannot be empty")

	// Aerospike URL validation failed due to invalid Aerospike DB port
	ErrInvalidPort = errors.New("aerospike port must be within 0 to 65536")

	// Aerospike URL validation failed due to empty Aerospike namespace
	ErrEmptyNamespace = errors.New("aerospike namespace cannot be empty")

	// Aerospike URL validation failed due to invalid Aerospike namespace
	ErrInvalidNamespace = errors.New("aerospike namespace is invalid")
)
