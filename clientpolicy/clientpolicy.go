// Clientpolicy package.
// Aerospike client policy [aerospike.ClientPolicy] properties are parsed from [aerourl.AerospikeURL] here.
//
// [aerospike.ClientPolicy]: https://pkg.go.dev/github.com/aerospike/aerospike-client-go/v6#ClientPolicy
package clientpolicy

import (
	"strconv"
	"strings"
	"time"

	"github.com/aerospike/aerospike-client-go/v6"
	"github.com/tiptophelmet/aerospike-url/aerourl"
	"github.com/tiptophelmet/aerospike-url/factory"
)

// Parses [aerospike.ClientPolicy] properties from validated [aerourl.AerospikeURL]
// If URL query is not empty, properties will be parsed.
//
// [aerospike.ClientPolicy]: https://pkg.go.dev/github.com/aerospike/aerospike-client-go/v6#ClientPolicy
func Parse(aeroURL *aerourl.AerospikeURL, clientFactory *factory.AerospikeClientFactory) {
	if len(aeroURL.GetNetURL().Query()) == 0 {
		return
	}

	policy := aerospike.NewClientPolicy()
	parser := &ClientPolicyParser{aeroURL, policy}

	parser.AuthMode()
	parser.User()
	parser.Password()
	parser.ClusterName()
	parser.Timeout()
	parser.IdleTimeout()
	parser.LoginTimeout()
	parser.ConnectionQueueSize()
	parser.MinConnectionsPerNode()
	parser.MaxErrorRate()
	parser.ErrorRateWindow()
	parser.LimitConnectionsToQueueSize()
	parser.OpeningConnectionThreshold()
	parser.FailIfNotConnected()
	parser.TendInterval()
	parser.UseServicesAlternate()
	parser.RackAware()
	parser.RackId()
	parser.IgnoreOtherSubnetAliases()
	parser.SeedOnlyCluster()

	clientFactory.SetClientPolicy(parser.GetClientPolicy())
}

// Serves as a holder for [aerourl.AerospikeURL] and [aerospike.ClientPolicy].
// Has a collection of methods to identify and parse each client policy property from URL query.
//
// [aerospike.ClientPolicy]: https://pkg.go.dev/github.com/aerospike/aerospike-client-go/v6#ClientPolicy
type ClientPolicyParser struct {
	aeroURL *aerourl.AerospikeURL
	policy  *aerospike.ClientPolicy
}

// Parses `aerospike.ClientPolicy.AuthMode`.
// See: https://pkg.go.dev/github.com/aerospike/aerospike-client-go/v6#ClientPolicy
func (parser *ClientPolicyParser) AuthMode() {
	authMode := parser.aeroURL.GetNetURL().Query().Get("auth_mode")

	if authMode == "auth_mode_internal" {
		parser.policy.AuthMode = aerospike.AuthModeInternal
	} else if authMode == "auth_mode_external" {
		parser.policy.AuthMode = aerospike.AuthModeExternal
	} else if authMode == "auth_mode_pki" {
		parser.policy.AuthMode = aerospike.AuthModePKI
	}
}

// Parses `aerospike.ClientPolicy.User`.
// See: https://pkg.go.dev/github.com/aerospike/aerospike-client-go/v6#ClientPolicy
func (parser *ClientPolicyParser) User() {
	user := parser.aeroURL.GetNetURL().User.Username()
	if user != "" {
		parser.policy.User = user
	}
}

// Parses `aerospike.ClientPolicy.Password`.
// See: https://pkg.go.dev/github.com/aerospike/aerospike-client-go/v6#ClientPolicy
func (parser *ClientPolicyParser) Password() {
	password, isPasswordSet := parser.aeroURL.GetNetURL().User.Password()
	if isPasswordSet {
		parser.policy.Password = password
	}
}

// Parses `aerospike.ClientPolicy.ClusterName`.
// See: https://pkg.go.dev/github.com/aerospike/aerospike-client-go/v6#ClientPolicy
func (parser *ClientPolicyParser) ClusterName() {
	clusterName := parser.aeroURL.GetNetURL().Query().Get("cluster_name")
	if clusterName != "" {
		parser.policy.ClusterName = clusterName
	}
}

// Parses `aerospike.ClientPolicy.Timeout`.
// See: https://pkg.go.dev/github.com/aerospike/aerospike-client-go/v6#ClientPolicy
func (parser *ClientPolicyParser) Timeout() {
	timeoutStr := parser.aeroURL.GetNetURL().Query().Get("timeout")

	if timeoutStr != "" {
		timeout, _ := time.ParseDuration(timeoutStr)
		parser.policy.Timeout = timeout
	}
}

// Parses `aerospike.ClientPolicy.IdleTimeout`.
// See: https://pkg.go.dev/github.com/aerospike/aerospike-client-go/v6#ClientPolicy
func (parser *ClientPolicyParser) IdleTimeout() {
	idleTimeoutStr := parser.aeroURL.GetNetURL().Query().Get("idle_timeout")

	if idleTimeoutStr != "" {
		idleTimeout, _ := time.ParseDuration(idleTimeoutStr)
		parser.policy.IdleTimeout = idleTimeout
	}
}

// Parses `aerospike.ClientPolicy.LoginTimeout`.
// See: https://pkg.go.dev/github.com/aerospike/aerospike-client-go/v6#ClientPolicy
func (parser *ClientPolicyParser) LoginTimeout() {
	loginTimeoutStr := parser.aeroURL.GetNetURL().Query().Get("login_timeout")

	if loginTimeoutStr != "" {
		loginTimeout, _ := time.ParseDuration(loginTimeoutStr)
		parser.policy.LoginTimeout = loginTimeout
	}
}

// Parses `aerospike.ClientPolicy.ConnectionQueueSize`.
// See: https://pkg.go.dev/github.com/aerospike/aerospike-client-go/v6#ClientPolicy
func (parser *ClientPolicyParser) ConnectionQueueSize() {
	connQueueSizeStr := strings.TrimSpace(parser.aeroURL.GetNetURL().Query().Get("connection_queue_size"))

	if connQueueSizeStr != "" {
		connQueueSize, _ := strconv.Atoi(connQueueSizeStr)
		parser.policy.ConnectionQueueSize = connQueueSize
	}
}

// Parses `aerospike.ClientPolicy.MinConnectionsPerNode`.
// See: https://pkg.go.dev/github.com/aerospike/aerospike-client-go/v6#ClientPolicy
func (parser *ClientPolicyParser) MinConnectionsPerNode() {
	minConnsPerNodeStr := strings.TrimSpace(parser.aeroURL.GetNetURL().Query().Get("min_connections_per_node"))

	if minConnsPerNodeStr != "" {
		minConnsPerNode, _ := strconv.Atoi(minConnsPerNodeStr)
		parser.policy.MinConnectionsPerNode = minConnsPerNode
	}
}

// Parses `aerospike.ClientPolicy.MaxErrorRate`.
// See: https://pkg.go.dev/github.com/aerospike/aerospike-client-go/v6#ClientPolicy
func (parser *ClientPolicyParser) MaxErrorRate() {
	maxErrorRateStr := strings.TrimSpace(parser.aeroURL.GetNetURL().Query().Get("max_error_rate"))

	if maxErrorRateStr != "" {
		maxErrorRate, _ := strconv.Atoi(maxErrorRateStr)
		parser.policy.MaxErrorRate = maxErrorRate
	}
}

// Parses `aerospike.ClientPolicy.ErrorRateWindow`.
// See: https://pkg.go.dev/github.com/aerospike/aerospike-client-go/v6#ClientPolicy
func (parser *ClientPolicyParser) ErrorRateWindow() {
	errorRateWindowStr := strings.TrimSpace(parser.aeroURL.GetNetURL().Query().Get("error_rate_window"))

	if errorRateWindowStr != "" {
		errorRateWindow, _ := strconv.Atoi(errorRateWindowStr)
		parser.policy.ErrorRateWindow = errorRateWindow
	}
}

// Parses `aerospike.ClientPolicy.LimitConnectionsToQueueSize`.
// See: https://pkg.go.dev/github.com/aerospike/aerospike-client-go/v6#ClientPolicy
func (parser *ClientPolicyParser) LimitConnectionsToQueueSize() {
	limitConnsToQueueSizeStr := strings.TrimSpace(parser.aeroURL.GetNetURL().Query().Get("limit_connections_to_queue_size"))

	if limitConnsToQueueSizeStr != "" {
		limitConnsToQueueSize, _ := strconv.ParseBool(limitConnsToQueueSizeStr)
		parser.policy.LimitConnectionsToQueueSize = limitConnsToQueueSize
	}
}

// Parses `aerospike.ClientPolicy.OpeningConnectionThreshold`.
// See: https://pkg.go.dev/github.com/aerospike/aerospike-client-go/v6#ClientPolicy
func (parser *ClientPolicyParser) OpeningConnectionThreshold() {
	openingConnThresholdStr := strings.TrimSpace(parser.aeroURL.GetNetURL().Query().Get("opening_connection_threshold"))

	if openingConnThresholdStr != "" {
		openingConnThreshold, _ := strconv.Atoi(openingConnThresholdStr)
		parser.policy.OpeningConnectionThreshold = openingConnThreshold
	}
}

// Parses `aerospike.ClientPolicy.FailIfNotConnected`.
// See: https://pkg.go.dev/github.com/aerospike/aerospike-client-go/v6#ClientPolicy
func (parser *ClientPolicyParser) FailIfNotConnected() {
	failIfNotConnectedStr := strings.TrimSpace(parser.aeroURL.GetNetURL().Query().Get("fail_if_not_connected"))

	if failIfNotConnectedStr != "" {
		failIfNotConnected, _ := strconv.ParseBool(failIfNotConnectedStr)
		parser.policy.FailIfNotConnected = failIfNotConnected
	}
}

// Parses `aerospike.ClientPolicy.TendInterval`.
// See: https://pkg.go.dev/github.com/aerospike/aerospike-client-go/v6#ClientPolicy
func (parser *ClientPolicyParser) TendInterval() {
	tendIntervalStr := parser.aeroURL.GetNetURL().Query().Get("tend_interval")

	if tendIntervalStr != "" {
		tendInterval, _ := time.ParseDuration(tendIntervalStr)
		parser.policy.TendInterval = tendInterval
	}
}

// Parses `aerospike.ClientPolicy.UseServicesAlternate`.
// See: https://pkg.go.dev/github.com/aerospike/aerospike-client-go/v6#ClientPolicy
func (parser *ClientPolicyParser) UseServicesAlternate() {
	useServicesAlternateStr := strings.TrimSpace(parser.aeroURL.GetNetURL().Query().Get("use_services_alternate"))

	if useServicesAlternateStr != "" {
		useServicesAlternate, _ := strconv.ParseBool(useServicesAlternateStr)
		parser.policy.UseServicesAlternate = useServicesAlternate
	}
}

// Parses `aerospike.ClientPolicy.RackAware`.
// See: https://pkg.go.dev/github.com/aerospike/aerospike-client-go/v6#ClientPolicy
func (parser *ClientPolicyParser) RackAware() {
	rackAwareStr := strings.TrimSpace(parser.aeroURL.GetNetURL().Query().Get("rack_aware"))

	if rackAwareStr != "" {
		rackAware, _ := strconv.ParseBool(rackAwareStr)
		parser.policy.RackAware = rackAware
	}
}

// Parses `aerospike.ClientPolicy.RackId`.
// See: https://pkg.go.dev/github.com/aerospike/aerospike-client-go/v6#ClientPolicy
func (parser *ClientPolicyParser) RackId() {
	rackIdStr := strings.TrimSpace(parser.aeroURL.GetNetURL().Query().Get("rack_id"))

	if rackIdStr != "" {
		rackId, _ := strconv.Atoi(rackIdStr)
		parser.policy.RackId = rackId
	}
}

// Parses `aerospike.ClientPolicy.IgnoreOtherSubnetAliases`.
// See: https://pkg.go.dev/github.com/aerospike/aerospike-client-go/v6#ClientPolicy
func (parser *ClientPolicyParser) IgnoreOtherSubnetAliases() {
	ignoreOtherSubnetAliasesStr := strings.TrimSpace(parser.aeroURL.GetNetURL().Query().Get("ignore_subnet_aliases"))

	if ignoreOtherSubnetAliasesStr != "" {
		ignoreOtherSubnetAliases, _ := strconv.ParseBool(ignoreOtherSubnetAliasesStr)
		parser.policy.IgnoreOtherSubnetAliases = ignoreOtherSubnetAliases
	}
}

// Parses `aerospike.ClientPolicy.SeedOnlyCluster`.
// See: https://pkg.go.dev/github.com/aerospike/aerospike-client-go/v6#ClientPolicy
func (parser *ClientPolicyParser) SeedOnlyCluster() {
	seedOnlyClusterStr := strings.TrimSpace(parser.aeroURL.GetNetURL().Query().Get("seed_only_cluster"))

	if seedOnlyClusterStr != "" {
		seedOnlyCluster, _ := strconv.ParseBool(seedOnlyClusterStr)
		parser.policy.SeedOnlyCluster = seedOnlyCluster
	}
}

// Returns [aerospike.ClientPolicy].
// [aerospike.ClientPolicy] https://pkg.go.dev/github.com/aerospike/aerospike-client-go/v6#ClientPolicy
func (parser *ClientPolicyParser) GetClientPolicy() *aerospike.ClientPolicy {
	return parser.policy
}
