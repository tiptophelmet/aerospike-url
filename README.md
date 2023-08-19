# aerospike-url
💫 Parses URL connection string into Aerospike client.

# 👇 Usage

### 🔗 Connection string format
`aerospike://user:password@localhost:6789?auth_mode=auth_mode_internal&cluster_name=aerocluster-001&timeout=30`

### ⚙️ Parse connection string into Aerospike client factory

```
package main

import (
    "github.com/tiptophelmet/aerospike-url"
    aero "github.com/aerospike/aerospike-client-go/v6"
)

// Use aerospikeurl to build Aerospike client from URL
func getClientFromURL(url string) {
    clientFactory, err := aerospikeurl.Parse(connStr)
    if err != nil {
        panic(err)
    }

    return clientFactory.GetClient()
}

// Use Aerospike client as usual
func main() {
    url := "aerospike://user:password@localhost:6789?auth_mode=auth_mode_internal&cluster_name=aerocluster-001&timeout=30"

    client := getClientFromURL(url)

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

    // delete the key, and check if key exists
    existed, err := client.Delete(nil, key)
    panicOnError(err)

    fmt.Printf("Record existed before delete? %v\n", existed)
}
```

# ⚠️ Limitations

The following client policy fields are not supported as URL query parameters & can be set by directly modifying ClientPolicy (to get it: `clientFactory.GetClientPolicy()`)
- `aerospike.ClientPolicy.IpMap`
- `aerospike.ClientPolicy.RackIds`
- `aerospike.ClientPolicy.TlsConfig`