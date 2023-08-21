package aerourl

import (
	"net/url"
	"strings"
)

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
