//go:build unit

package aerospikeurl

import (
	"testing"

	"github.com/tiptophelmet/aerospike-url/aerourl"
)

func TestGenerateClientFactoryWithInvalidAeroURL(t *testing.T) {
	aeroURL, _ := aerourl.Init("https://127.0.0.1:")

	clientFactory, err := generateClientFactory(aeroURL)

	if err == nil {
		t.Errorf("got: %v, want error != nil", err)
	}

	if clientFactory != nil {
		t.Errorf("got: %v, want: *factory.ClientFactory = nil", clientFactory)
	}
}

func TestGenerateClientFactory(t *testing.T) {
	aeroURL, _ := aerourl.Init("aerospike://127.0.0.1:3000/aero-namespace-001")

	clientFactory, err := generateClientFactory(aeroURL)

	if err != nil {
		t.Errorf("got: %v, want error = nil", err)
	}

	if clientFactory == nil {
		t.Errorf("got: %v, want: *factory.ClientFactory != nil", clientFactory)
	}
}

func TestParseWithInvalidConnectionString(t *testing.T) {
	clientFactory, err := Parse("https://127.0.0.1:")

	if err == nil {
		t.Errorf("got: %v, want error != nil", err)
	}

	if clientFactory != nil {
		t.Errorf("got: %v, want: *factory.ClientFactory = nil", clientFactory)
	}
}

func TestParse(t *testing.T) {
	clientFactory, err := Parse("aerospike://127.0.0.1:3000/aero-namespace-001")

	if err != nil {
		t.Errorf("got: %v, want error = nil", err)
	}

	if clientFactory == nil {
		t.Errorf("got: %v, want: *factory.ClientFactory != nil", clientFactory)
	}
}
