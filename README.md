# aerospike-url 
[![GitHub version](https://badge.fury.io/gh/tiptophelmet%2Faerospike-url.svg)](https://badge.fury.io/gh/tiptophelmet%2Faerospike-url)
[![Go project version](https://badge.fury.io/go/github.com%2Ftiptophelmet%2Faerospike-url.svg)](https://badge.fury.io/go/github.com%2Ftiptophelmet%2Faerospike-url)
[![codecov](https://codecov.io/gh/tiptophelmet/aerospike-url/graph/badge.svg?token=PDE9SG7H4Y)](https://codecov.io/gh/tiptophelmet/aerospike-url)

💫 Parses URL connection string into Aerospike client.

# 👇 Usage

### 🔗 Connection string format
`aerospike://aerouser001:aerouserpassw123@127.0.0.1:3000/my-aerospike-namespace?auth_mode=auth_mode_internal&timeout=10s&idle_timeout=3s&max_error_rate=50`

### ⚙️ Parse connection string into Aerospike client factory

```
package main

import (
	"fmt"

	aerospikeurl "github.com/tiptophelmet/aerospike-url"
	aero "github.com/aerospike/aerospike-client-go/v6"
)

// Use aerospikeurl to build Aerospike client from URL
func buildClientFromURL(url string) *aero.Client {
	clientFactory, err := aerospikeurl.Parse(url)
	if err != nil {
		panic(err)
	}

	client, err := clientFactory.BuildClient()
	if err != nil {
		panic(err)
	}

	return client
}

// This is only for this example.
// Please handle errors properly.
func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

// Use Aerospike client as usual
func main() {
	url := "aerospike://aerouser001:aerouserpassw123@127.0.0.1:3000/my-aerospike-namespace?auth_mode=auth_mode_internal&timeout=10s&idle_timeout=3s&max_error_rate=50"

	client := buildClientFromURL(url)
	defer client.Close()

	key, err := aero.NewKey("test", "aerospike", "key")
	panicOnError(err)

	// define some bins with data
	bins := aero.BinMap{
		"bin1": 42,
		"bin2": "An elephant is a mouse with an operating system",
		"bin3": []interface{}{"Go", 2009},
	}

	// write the bins
	err = client.Put(nil, key, bins)
	panicOnError(err)

	// read it back!
	rec, err := client.Get(nil, key)
	panicOnError(err)

	fmt.Printf("Record bins: %v", rec.Bins)

	// delete the key, and check if key exists
	existed, err := client.Delete(nil, key)
	panicOnError(err)

	fmt.Printf("Record existed before delete? %v\n", existed)
}
```

# ⚠️ Limitations

1. The following client policy fields are not supported as URL query parameters & can be set by directly modifying ClientPolicy (to get it: `clientFactory.GetClientPolicy()`)
- `aerospike.ClientPolicy.IpMap`
- `aerospike.ClientPolicy.RackIds`
- `aerospike.ClientPolicy.TlsConfig`

---

2. `AerospikeClientFactory.GetClient()` uses the following methods for client generation:
- `aerospike.NewClient(hostname string, port int) (*Client, Error)`
- `aerospike.NewClientWithPolicy(policy *ClientPolicy, hostname string, port int) (*Client, Error)`


`AerospikeClientFactory` does NOT support the following method as `hosts` cannot be specified in aerospike URL:
- `aerospike.NewClientWithPolicyAndHost(policy *ClientPolicy, hosts ...*Host) (*Client, Error)`