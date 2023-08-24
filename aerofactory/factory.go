// Factory package.
// Aerospike DB hostname, port & client policy are assembled into DB client factory here.
package aerofactory

import (
	"github.com/aerospike/aerospike-client-go/v6"
)

// Serves for assembling hostname, port & [aerospike.ClientPolicy] for building [aerospike.Client].
//
// [aerospike.ClientPolicy]: https://pkg.go.dev/github.com/aerospike/aerospike-client-go/v6#ClientPolicy
// [aerospike.Client]: https://pkg.go.dev/github.com/aerospike/aerospike-client-go/v6#Client
type AerospikeClientFactory struct {
	hostname  string
	port      int
	namespace string

	policy *aerospike.ClientPolicy
}

func (cf *AerospikeClientFactory) SetAddress(hostname string, port int, namespace string) {
	cf.hostname = hostname
	cf.port = port
	cf.namespace = namespace
}

// Sets Aerospike DB hostname.
func (cf *AerospikeClientFactory) SetHostname(hostname string) {
	cf.hostname = hostname
}

// Returns Aerospike DB hostname.
func (cf *AerospikeClientFactory) GetHostname() string {
	return cf.hostname
}

// Sets Aerospike DB port.
func (cf *AerospikeClientFactory) SetPort(port int) {
	cf.port = port
}

// Returns Aerospike DB port.
func (cf *AerospikeClientFactory) GetPort() int {
	return cf.port
}

// Sets Aerospike DB namespace.
func (cf *AerospikeClientFactory) SetNamespace(namespace string) {
	cf.namespace = namespace
}

// Returns Aerospike DB namespace.
func (cf *AerospikeClientFactory) GetNamespace() string {
	return cf.namespace
}

// Sets Aerospike DB client policy [aerospike.ClientPolicy].
//
// [aerospike.ClientPolicy]: https://pkg.go.dev/github.com/aerospike/aerospike-client-go/v6#ClientPolicy
func (cf *AerospikeClientFactory) SetClientPolicy(policy *aerospike.ClientPolicy) {
	cf.policy = policy
}

// Returns Aerospike DB client policy [aerospike.ClientPolicy].
//
// [aerospike.ClientPolicy]: https://pkg.go.dev/github.com/aerospike/aerospike-client-go/v6#ClientPolicy
func (cf *AerospikeClientFactory) GetClientPolicy() *aerospike.ClientPolicy {
	return cf.policy
}

// Builds Aerospike DB client [aerospike.Client].
// If [aerospike.ClientPolicy] was parsed from [aerourl.AerospikeURL],
// client is created using [aerospike.NewClientWithPolicy], otherwise it is created using [aerospike.NewClient].
//
// [aerospike.Client]: https://pkg.go.dev/github.com/aerospike/aerospike-client-go/v6#Client
// [aerospike.ClientPolicy]: https://pkg.go.dev/github.com/aerospike/aerospike-client-go/v6#ClientPolicy
// [aerospike.NewClientWithPolicy]: https://pkg.go.dev/github.com/aerospike/aerospike-client-go/v6#NewClientWithPolicy
// [aerospike.NewClient]: https://pkg.go.dev/github.com/aerospike/aerospike-client-go/v6#NewClient
func (cf *AerospikeClientFactory) BuildClient() (*aerospike.Client, aerospike.Error) {
	if cf.policy == nil {
		return aerospike.NewClient(cf.hostname, cf.port)
	}

	return aerospike.NewClientWithPolicy(cf.policy, cf.hostname, cf.port)
}
