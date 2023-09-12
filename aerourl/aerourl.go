// Aerospike URL package.
// Aerospike connection string is resolved & processed into [aerourl.AerospikeURL] here.
package aerourl

import (
	"net/url"
	"strconv"
	"strings"
)

// Processes connection string, initializes, resolves & returns [aerourl.AerospikeURL].
// Returns error, if connection string is empty or not resolved.
func Init(connStr string) (*AerospikeURL, error) {
	if strings.TrimSpace(connStr) == "" {
		return nil, ErrEmptyConnStr
	}

	connURL, err := url.Parse(connStr)
	if err != nil {
		return nil, err
	}

	aeroURL := &AerospikeURL{url: connURL}

	err = aeroURL.resolve()
	if err != nil {
		return nil, err
	}

	return aeroURL, nil
}

// Serves as a validated [net/url.URL] that possesses min required data for creating Aerospike DB client.
// It is strictly instantiated using [aerourl.Init] to get a properly validated Aerospike URL.
//
// Instantiating this struct directly will result in failure to get min required data for DB client.
type AerospikeURL struct {
	hostname  string
	port      int
	namespace string

	url *url.URL
}

// Resolves min required data required for creating Aerospike DB client.
// If URL is missing a required part - error is returned, otherwise it is nil.
func (aeroURL *AerospikeURL) resolve() error {
	if aeroURL.url.Scheme != "aerospike" {
		return ErrInvalidScheme
	}

	if err := aeroURL.resolveHostname(); err != nil {
		return err
	}

	if err := aeroURL.resolvePort(); err != nil {
		return err
	}

	if err := aeroURL.resolveNamespace(); err != nil {
		return err
	}

	return nil
}

// Validates and sets Aerospike hostname.
// If hostname is empty, error is returned.
func (aeroURL *AerospikeURL) resolveHostname() error {
	if aeroURL.url.Hostname() == "" {
		return ErrEmptyHostname
	}

	aeroURL.hostname = aeroURL.url.Hostname()
	return nil
}

// Validates and sets Aerospike port.
// If port is empty or invalid, error is returned.
func (aeroURL *AerospikeURL) resolvePort() error {
	if aeroURL.url.Port() == "" {
		return ErrEmptyPort
	}

	port, _ := strconv.Atoi(aeroURL.url.Port())

	aeroURL.port = port
	return nil
}

// Validates and sets [Aerospike namespace].
// If port is empty or invalid, error is returned.
//
// [Aerospike namespace] https://docs.aerospike.com/server/architecture/data-model#namespaces
func (aeroURL *AerospikeURL) resolveNamespace() error {
	pathParts := strings.Split(aeroURL.url.Path, "/")
	if len(pathParts) == 1 {
		return ErrEmptyNamespace
	}

	namespace := ""

	if len(pathParts) >= 2 {
		namespace = strings.TrimSpace(pathParts[1])
		if len(namespace) == 0 {
			return ErrEmptyNamespace
		}

		if strings.ContainsAny(namespace, " \t\n\r") {
			return ErrInvalidNamespace
		}
	}

	aeroURL.namespace = namespace
	return nil
}

// Returns resolved Aerospike hostname.
func (aerourl *AerospikeURL) Hostname() string {
	return aerourl.hostname
}

// Returns resolved Aerospike hostname.
func (aerourl *AerospikeURL) Port() int {
	return aerourl.port
}

// Returns resolved [Aerospike namespace].
//
// [Aerospike namespace] https://docs.aerospike.com/server/architecture/data-model#namespaces
func (aerourl *AerospikeURL) Namespace() string {
	return aerourl.namespace
}

// Retrieves underlying [net/url.URL].
// Empty [net/url.URL] is returned, if it was not initialized & validated at [aerourl.Init].
//
// Empty or non-validated url will result in failure to get min required data for DB client.
func (aeroURL *AerospikeURL) GetNetURL() *url.URL {
	if aeroURL.url == nil {
		aeroURL.url = &url.URL{}
	}

	return aeroURL.url
}
