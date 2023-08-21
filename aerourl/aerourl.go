package aerourl

import (
	"errors"
	"net/url"
)

func Init(connStr string) (*AerospikeURL, error) {
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

func (aeroURL *AerospikeURL) GetURL() *url.URL {
	return aeroURL.url
}

var (
	ErrNilURL        = errors.New("aerospike URL must be initialized with connection string")
	ErrInvalidScheme = errors.New("invalid url scheme, want: aerospike://")
	ErrEmptyHostname = errors.New("aerospike hostname cannot be empty")
	ErrEmptyPort     = errors.New("aerospike port cannot be empty")
)
