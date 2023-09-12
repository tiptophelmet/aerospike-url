package aerourl

import (
	"errors"
	"fmt"
	"net/url"
	"testing"
)

func TestInitEmptyConnStr(t *testing.T) {
	connStr := ""
	aeroURL, err := Init(connStr)

	if !errors.Is(err, ErrEmptyConnStr) {
		t.Fatalf(`got: %v, want: error is aerourl.ErrEmptyConnStr`, err)
	}

	if aeroURL != nil {
		t.Fatalf(`got: %v, want: *AerospikeURL == nil`, aeroURL)
	}
}

func TestInitEmptyScheme(t *testing.T) {
	connStr := "127.0.0.1:3000/aero-namespace-001"
	aeroURL, err := Init(connStr)

	var urlErr *url.Error
	if !errors.As(err, &urlErr) {
		t.Fatalf(`got: %v, want: error is *url.Error`, err)
	}

	if aeroURL != nil {
		t.Fatalf(`got: %v, want: *AerospikeURL == nil`, aeroURL)
	}
}

func TestInitEmptySchemeWithForwardSlashes(t *testing.T) {
	connStr := "//127.0.0.1:3000/aero-namespace-001"
	aeroURL, err := Init(connStr)

	if !errors.Is(err, ErrInvalidScheme) {
		t.Fatalf(`got: %v, want: error is aerourl.ErrInvalidScheme`, err)
	}

	if aeroURL != nil {
		t.Fatalf(`got: %v, want: *AerospikeURL == nil`, aeroURL)
	}
}

func TestInitInvalidScheme(t *testing.T) {
	connStr := "https://127.0.0.1:3000/aero-namespace-001"
	aeroURL, err := Init(connStr)

	if !errors.Is(err, ErrInvalidScheme) {
		t.Fatalf(`got: %v, want: error is aerourl.ErrInvalidScheme`, err)
	}

	if aeroURL != nil {
		t.Fatalf(`got: %v, want: *AerospikeURL == nil`, aeroURL)
	}
}

func TestInitEmptyHostname(t *testing.T) {
	connStr := "aerospike://:3000/aero-namespace-001"
	aeroURL, err := Init(connStr)

	if !errors.Is(err, ErrEmptyHostname) {
		t.Fatalf(`got: %v, want: error is aerourl.ErrEmptyHostname`, aeroURL)
	}

	if aeroURL != nil {
		t.Fatalf(`got: %v, want: *AerospikeURL == nil`, err)
	}
}

func TestInitEmptyPort(t *testing.T) {
	connStr := "aerospike://127.0.0.1:/aero-namespace-001"
	aeroURL, err := Init(connStr)

	if !errors.Is(err, ErrEmptyPort) {
		t.Fatalf(`got: %v, want: error is aerourl.ErrEmptyPort`, err)
	}

	if aeroURL != nil {
		t.Fatalf(`got: %v, want: *AerospikeURL == nil`, aeroURL)
	}
}

func TestInitEmptyNamespace(t *testing.T) {
	connStr := "aerospike://127.0.0.1:3000"
	aeroURL, err := Init(connStr)

	if !errors.Is(err, ErrEmptyNamespace) {
		t.Fatalf(`got: %v, want: error is aerourl.ErrEmptyNamespace`, err)
	}

	if aeroURL != nil {
		t.Fatalf(`got: %v, want: *AerospikeURL == nil`, aeroURL)
	}
}

func TestInitInvalidNamespace(t *testing.T) {
	connStr := "aerospike://127.0.0.1:3000/ my aero namespace "
	aeroURL, err := Init(connStr)

	if !errors.Is(err, ErrInvalidNamespace) {
		fmt.Println(aeroURL.GetNetURL().RawFragment)
		t.Fatalf(`got: %v, want: error is aerourl.ErrInvalidNamespace`, err)
	}

	if aeroURL != nil {
		t.Fatalf(`got: %v, want: *AerospikeURL == nil`, aeroURL)
	}
}

func TestInitNamespaceFromCompatiblePath(t *testing.T) {
	connStr := "aerospike://127.0.0.1:3000/aero-namespace-001/random/compat/path"
	aeroURL, err := Init(connStr)

	if err != nil {
		t.Fatalf(`got: %v, want: error = nil`, err)
	}

	if aeroURL == nil {
		t.Fatalf(`got: %v, want: *AerospikeURL != nil`, aeroURL)
	}
}

func TestInit(t *testing.T) {
	connStr := "aerospike://127.0.0.1:3000/aero-namespace-001"
	aeroURL, err := Init(connStr)

	if err != nil {
		t.Fatalf(`got: %v, want: error = nil`, err)
	}

	if aeroURL == nil {
		t.Fatalf(`got: %v, want: *AerospikeURL != nil`, aeroURL)
	}
}

func TestGetURL(t *testing.T) {
	connStr := "aerospike://127.0.0.1:3000/aero-namespace-001"
	aeroURL, err := Init(connStr)

	if err != nil {
		t.Fatalf(`got: %v, want: error = nil`, err)
	}

	if aeroURL == nil {
		t.Fatalf(`got: %v, want: *AerospikeURL != nil`, aeroURL)
	}

	connURL := aeroURL.GetNetURL()
	if connURL == nil {
		t.Fatal("got: *url.URL = nil, want *url.URL != nil")
	}
}
