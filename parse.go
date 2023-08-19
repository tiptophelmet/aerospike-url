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

	client := &AerospikeClientFactory{}

	port, _ := strconv.Atoi(connURL.Port())

	client.SetHostname(connURL.Hostname())
	client.SetPort(port)

	parseClientPolicyQuery(connURL, client)

	return nil, nil
}

func validateSchemeAndHost(connURL *url.URL) error {
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
