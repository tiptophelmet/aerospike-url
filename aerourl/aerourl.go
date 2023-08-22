// Aerospike URL package.
// Aerospike connection string is validated & processed into [aerourl.AerospikeURL] here.
package aerourl

import (
	"net/url"
	"strings"
)

// Processes connection string, initializes, validates & returns [aerourl.AerospikeURL].
// Returns error, if connection string is empty or failed validation.
func Init(connStr string) (*AerospikeURL, error) {
	if strings.TrimSpace(connStr) == "" {
		return nil, ErrEmptyConnStr
	}

	connURL, err := url.Parse(connStr)
	if err != nil {
		return nil, err
	}

	aeroURL := &AerospikeURL{connURL}

	err = aeroURL.validateSchemeAndHost()
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
	url *url.URL
}

func (aeroURL *AerospikeURL) validateSchemeAndHost() error {
	if aeroURL.url == nil {
		return ErrNilURL
	}

	if aeroURL.url.Scheme != "aerospike" {
		return ErrInvalidScheme
	}

	if aeroURL.url.Hostname() == "" {
		return ErrEmptyHostname
	}

	if aeroURL.url.Port() == "" {
		return ErrEmptyPort
	}

	return nil
}

// Retrieves underlying [net/url.URL].
// Empty [net/url.URL] is returned, if it was not initialized & validated at [aerourl.Init].
//
// Empty or non-validated url will result in failure to get min required data for DB client.
func (aeroURL *AerospikeURL) GetURL() *url.URL {
	if aeroURL.url == nil {
		aeroURL.url = &url.URL{}
	}

	return aeroURL.url
}
