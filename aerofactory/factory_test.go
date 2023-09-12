package aerofactory

import (
	"testing"

	"github.com/aerospike/aerospike-client-go/v6"
)

func TestSetAddress(t *testing.T) {
	factory := &AerospikeClientFactory{}

	var (
		hostname  string = "127.0.0.1"
		port      int    = 3000
		namespace string = "aero-namespace-001"
	)

	factory.SetAddress(hostname, port, namespace)

	gotHostname := factory.GetHostname()
	if gotHostname != hostname {
		t.Errorf("got: %v, want: %v", gotHostname, hostname)
	}

	gotPort := factory.GetPort()
	if gotPort != port {
		t.Errorf("got: %v, want: %v", gotPort, port)
	}

	gotNamespace := factory.GetNamespace()
	if gotNamespace != namespace {
		t.Errorf("got: %v, want: %v", gotNamespace, namespace)
	}
}

func TestSetGetHostname(t *testing.T) {
	factory := &AerospikeClientFactory{}

	hostname := "127.0.0.1"

	factory.SetHostname(hostname)
	gotHostname := factory.GetHostname()

	if gotHostname != hostname {
		t.Errorf("got: %v, want: %v", gotHostname, hostname)
	}
}

func TestSetGetPort(t *testing.T) {
	factory := &AerospikeClientFactory{}

	port := 3000

	factory.SetPort(port)
	gotPort := factory.GetPort()

	if gotPort != port {
		t.Errorf("got: %v, want: %v", gotPort, port)
	}
}

func TestSetGetNamespace(t *testing.T) {
	factory := &AerospikeClientFactory{}

	namespace := "aero-namespace-001"

	factory.SetNamespace(namespace)
	gotNamespace := factory.GetNamespace()

	if gotNamespace != namespace {
		t.Errorf("got: %v, want: %v", gotNamespace, namespace)
	}
}

func TestSetGetClientPolicy(t *testing.T) {
	factory := &AerospikeClientFactory{}

	policy := aerospike.NewClientPolicy()
	policy.AuthMode = aerospike.AuthModeInternal
	policy.User = "aero-user-001"
	policy.Password = "aerouser001passw"

	factory.SetClientPolicy(policy)
	gotPolicy := factory.GetClientPolicy()

	if gotPolicy.AuthMode != policy.AuthMode ||
		gotPolicy.User != policy.User ||
		gotPolicy.Password != policy.Password {
		t.Errorf("got: %v, want: %v", gotPolicy, policy)
	}
}
