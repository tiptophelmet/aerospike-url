package aerospikeurl

import (
	"errors"
	"strconv"

	"github.com/tiptophelmet/aerospike-url/aerourl"
	"github.com/tiptophelmet/aerospike-url/clientpolicy"
	"github.com/tiptophelmet/aerospike-url/factory"
)

func Parse(connStr string) (*factory.AerospikeClientFactory, error) {
	aeroURL, err := aerourl.Init(connStr)
	if err != nil {
		return nil, err
	}

	return generateClientFactory(aeroURL)
}

func generateClientFactory(aeroURL *aerourl.AerospikeURL) (*factory.AerospikeClientFactory, error) {
	if aeroURL == nil {
		return nil, errors.New("connURL must be initialized with connection string")
	}

	clientFactory := &factory.AerospikeClientFactory{}
	url := aeroURL.GetURL()
	port, _ := strconv.Atoi(url.Port())

	clientFactory.SetHostname(url.Hostname())
	clientFactory.SetPort(port)

	clientpolicy.Parse(aeroURL, clientFactory)

	return clientFactory, nil
}
