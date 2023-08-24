// Root package
package aerospikeurl

import (
	"errors"

	"github.com/tiptophelmet/aerospike-url/aerofactory"
	"github.com/tiptophelmet/aerospike-url/aerourl"
	"github.com/tiptophelmet/aerospike-url/clientpolicy"
)

// Parses Aerospike connection string into [factory.AerospikeClientFactory].
//
// Connection string format:
// ```aerospike://aero-user-001:aerouserpassw123@127.0.0.1:3000?auth_mode=auth_mode_internal&timeout=30```
func Parse(connStr string) (*aerofactory.AerospikeClientFactory, error) {
	aeroURL, err := aerourl.Init(connStr)
	if err != nil {
		return nil, err
	}

	return generateClientFactory(aeroURL)
}

// Generates [factory.AerospikeClientFactory] based on validated Aerospike URL [aerourl.AerospikeURL].
// Retrieves Aerospike DB hostname, port & client policy (See: [clientpolicy.Parse] and [aerospike.ClientPolicy]),
// then puts them into a client factory.
//
// [aerospike.ClientPolicy]: https://pkg.go.dev/github.com/aerospike/aerospike-client-go/v6
func generateClientFactory(aeroURL *aerourl.AerospikeURL) (*aerofactory.AerospikeClientFactory, error) {
	if aeroURL == nil {
		return nil, errors.New("connURL must be initialized with connection string")
	}

	clientFactory := &aerofactory.AerospikeClientFactory{}

	clientFactory.SetAddress(aeroURL.Hostname(), aeroURL.Port(), aeroURL.Namespace())
	clientpolicy.Parse(aeroURL, clientFactory)

	return clientFactory, nil
}
