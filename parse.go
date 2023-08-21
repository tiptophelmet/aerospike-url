package aerospikeurl

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
)

func Parse(connStr string) (*AerospikeClientFactory, error) {
	connURL, err := url.Parse(connStr)
	if err != nil {
		return nil, err
	}

	err = validateSchemeAndHost(connURL)
	if err != nil {
		return nil, err
	}

	return generateClientFactory(connURL), nil
}

func validateSchemeAndHost(connURL *url.URL) error {
	if connURL == nil {
		return errors.New("connURL must be initialized with connection string")
	}

	if connURL.Scheme != "aerospike" {
		return fmt.Errorf("invalid scheme: %v://, expected: aerospike://", connURL.Scheme)
	}

	if connURL.Hostname() == "" {
		return errors.New("aerospike hostname cannot be empty")
	}

	if connURL.Port() == "" {
		return errors.New("aerospike port cannot be empty")
	}

	return nil
}

func generateClientFactory(connURL *url.URL) *AerospikeClientFactory {
	clientFactory := &AerospikeClientFactory{}

	port, _ := strconv.Atoi(connURL.Port())

	clientFactory.SetHostname(connURL.Hostname())
	clientFactory.SetPort(port)

	parseClientPolicyFields(connURL, clientFactory)

	return clientFactory
}
