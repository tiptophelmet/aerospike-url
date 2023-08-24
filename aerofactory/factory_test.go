package aerofactory

import (
	"flag"
	"os"
	"testing"
	"time"

	"github.com/aerospike/aerospike-client-go/v6"
)

var aerospikeHostname string
var aerospikePort int

func TestMain(m *testing.M) {
	flag.StringVar(&aerospikeHostname, "hostname", "", "Aerospike hostname")
	flag.IntVar(&aerospikePort, "port", 0, "Aerospike port")

	exitCode := m.Run()
	os.Exit(exitCode)
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

func TestBuildClient(t *testing.T) {
	if aerospikeHostname == "" || aerospikePort == 0 {
		t.Skip("set aerospike hostname & port to run this test (go test -hostname='127.0.0.1' -port='3000')")
	}

	factory := &AerospikeClientFactory{}

	factory.SetHostname(aerospikeHostname)
	factory.SetPort(aerospikePort)

	client, err := factory.BuildClient()
	if err != nil {
		t.Fatalf("got: %v, want error = nil", err)
	}

	if client == nil {
		t.Fatalf("got: %v, want *aerospike.Client != nil", client)
	}

	if !client.IsConnected() {
		t.Fatal("got: client.IsConnected() = false, want: client.IsConnected() = true")
	}

	client.Close()
}

func TestBuildClientWithPolicy(t *testing.T) {
	if aerospikeHostname == "" || aerospikePort == 0 {
		t.Skip("set aerospike hostname & port to run this test (go test -hostname='127.0.0.1' -port='3000')")
	}

	factory := &AerospikeClientFactory{}

	factory.SetHostname(aerospikeHostname)
	factory.SetPort(aerospikePort)

	policy := aerospike.NewClientPolicy()

	timeout, _ := time.ParseDuration("10s")
	idleTimeout, _ := time.ParseDuration("3s")

	policy.Timeout = timeout
	policy.IdleTimeout = idleTimeout
	policy.MaxErrorRate = 50

	factory.SetClientPolicy(policy)

	client, err := factory.BuildClient()
	if err != nil {
		t.Fatalf("got: %v, want error = nil", err)
	}

	if client == nil {
		t.Fatalf("got: %v, want *aerospike.Client != nil", client)
	}

	if !client.IsConnected() {
		t.Fatal("got: client.IsConnected() = false, want: client.IsConnected() = true")
	}

	client.Close()
}
