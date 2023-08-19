package aerospikeurl

import "github.com/aerospike/aerospike-client-go/v6"

type AerospikeClientFactory struct {
	policy   *aerospike.ClientPolicy
	hostname string
	port     int
}

func (cf *AerospikeClientFactory) SetHostname(hostname string) {
	cf.hostname = hostname
}

func (cf *AerospikeClientFactory) SetPort(port int) {
	cf.port = port
}

func (cf *AerospikeClientFactory) GetClientPolicy() *aerospike.ClientPolicy {
	if cf.policy == nil {
		cf.policy = aerospike.NewClientPolicy()
	}

	return cf.policy
}

func (cf *AerospikeClientFactory) GetClient() (*aerospike.Client, aerospike.Error) {
	if cf.policy == nil {
		return aerospike.NewClient(cf.hostname, cf.port)
	}

	return aerospike.NewClientWithPolicy(cf.policy, cf.hostname, cf.port)
}
