package clientpolicy

import (
	"strconv"
	"strings"
	"time"

	"github.com/aerospike/aerospike-client-go/v6"
	"github.com/tiptophelmet/aerospike-url/aerourl"
	"github.com/tiptophelmet/aerospike-url/factory"
)

func Parse(aeroURL *aerourl.AerospikeURL, client *factory.AerospikeClientFactory) {
	if len(aeroURL.GetURL().Query()) == 0 {
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

	client.SetClientPolicy(parser.GetClientPolicy())
}

type ClientPolicyParser struct {
	aeroURL *aerourl.AerospikeURL
	policy  *aerospike.ClientPolicy
}

func (parser *ClientPolicyParser) GetClientPolicy() *aerospike.ClientPolicy {
	return parser.policy
}

func (parser *ClientPolicyParser) AuthMode() {
	authMode := parser.aeroURL.GetURL().Query().Get("auth_mode")

	if authMode == "auth_mode_internal" {
		parser.policy.AuthMode = aerospike.AuthModeInternal
	} else if authMode == "auth_mode_external" {
		parser.policy.AuthMode = aerospike.AuthModeExternal
	} else if authMode == "auth_mode_pki" {
		parser.policy.AuthMode = aerospike.AuthModePKI
	}
}

func (parser *ClientPolicyParser) User() {
	user := parser.aeroURL.GetURL().User.Username()
	if user != "" {
		parser.policy.User = user
	}
}

func (parser *ClientPolicyParser) Password() {
	password, isPasswordSet := parser.aeroURL.GetURL().User.Password()
	if isPasswordSet {
		parser.policy.Password = password
	}
}

func (parser *ClientPolicyParser) ClusterName() {
	clusterName := parser.aeroURL.GetURL().Query().Get("cluster_name")
	if clusterName != "" {
		parser.policy.ClusterName = clusterName
	}
}

func (parser *ClientPolicyParser) Timeout() {
	timeoutStr := parser.aeroURL.GetURL().Query().Get("timeout")

	if timeoutStr != "" {
		timeout, _ := time.ParseDuration(timeoutStr)
		parser.policy.Timeout = timeout
	}
}

func (parser *ClientPolicyParser) IdleTimeout() {
	idleTimeoutStr := parser.aeroURL.GetURL().Query().Get("idle_timeout")

	if idleTimeoutStr != "" {
		idleTimeout, _ := time.ParseDuration(idleTimeoutStr)
		parser.policy.IdleTimeout = idleTimeout
	}
}

func (parser *ClientPolicyParser) LoginTimeout() {
	loginTimeoutStr := parser.aeroURL.GetURL().Query().Get("login_timeout")

	if loginTimeoutStr != "" {
		loginTimeout, _ := time.ParseDuration(loginTimeoutStr)
		parser.policy.LoginTimeout = loginTimeout
	}
}

func (parser *ClientPolicyParser) ConnectionQueueSize() {
	connQueueSizeStr := strings.TrimSpace(parser.aeroURL.GetURL().Query().Get("connection_queue_size"))

	if connQueueSizeStr != "" {
		connQueueSize, _ := strconv.Atoi(connQueueSizeStr)
		parser.policy.ConnectionQueueSize = connQueueSize
	}
}

func (parser *ClientPolicyParser) MinConnectionsPerNode() {
	minConnsPerNodeStr := strings.TrimSpace(parser.aeroURL.GetURL().Query().Get("min_connections_per_node"))

	if minConnsPerNodeStr != "" {
		minConnsPerNode, _ := strconv.Atoi(minConnsPerNodeStr)
		parser.policy.MinConnectionsPerNode = minConnsPerNode
	}
}

func (parser *ClientPolicyParser) MaxErrorRate() {
	maxErrorRateStr := strings.TrimSpace(parser.aeroURL.GetURL().Query().Get("max_error_rate"))

	if maxErrorRateStr != "" {
		maxErrorRate, _ := strconv.Atoi(maxErrorRateStr)
		parser.policy.MaxErrorRate = maxErrorRate
	}
}

func (parser *ClientPolicyParser) ErrorRateWindow() {
	errorRateWindowStr := strings.TrimSpace(parser.aeroURL.GetURL().Query().Get("error_rate_window"))

	if errorRateWindowStr != "" {
		errorRateWindow, _ := strconv.Atoi(errorRateWindowStr)
		parser.policy.ErrorRateWindow = errorRateWindow
	}
}

func (parser *ClientPolicyParser) LimitConnectionsToQueueSize() {
	limitConnsToQueueSizeStr := strings.TrimSpace(parser.aeroURL.GetURL().Query().Get("limit_connections_to_queue_size"))

	if limitConnsToQueueSizeStr != "" {
		limitConnsToQueueSize, _ := strconv.ParseBool(limitConnsToQueueSizeStr)
		parser.policy.LimitConnectionsToQueueSize = limitConnsToQueueSize
	}
}

func (parser *ClientPolicyParser) OpeningConnectionThreshold() {
	openingConnThresholdStr := strings.TrimSpace(parser.aeroURL.GetURL().Query().Get("opening_connection_threshold"))

	if openingConnThresholdStr != "" {
		openingConnThreshold, _ := strconv.Atoi(openingConnThresholdStr)
		parser.policy.OpeningConnectionThreshold = openingConnThreshold
	}
}

func (parser *ClientPolicyParser) FailIfNotConnected() {
	failIfNotConnectedStr := strings.TrimSpace(parser.aeroURL.GetURL().Query().Get("fail_if_not_connected"))

	if failIfNotConnectedStr != "" {
		failIfNotConnected, _ := strconv.ParseBool(failIfNotConnectedStr)
		parser.policy.FailIfNotConnected = failIfNotConnected
	}
}

func (parser *ClientPolicyParser) TendInterval() {
	tendIntervalStr := parser.aeroURL.GetURL().Query().Get("tend_interval")

	if tendIntervalStr != "" {
		tendInterval, _ := time.ParseDuration(tendIntervalStr)
		parser.policy.TendInterval = tendInterval
	}
}

func (parser *ClientPolicyParser) UseServicesAlternate() {
	useServicesAlternateStr := strings.TrimSpace(parser.aeroURL.GetURL().Query().Get("use_services_alternate"))

	if useServicesAlternateStr != "" {
		useServicesAlternate, _ := strconv.ParseBool(useServicesAlternateStr)
		parser.policy.UseServicesAlternate = useServicesAlternate
	}
}

func (parser *ClientPolicyParser) RackAware() {
	rackAwareStr := strings.TrimSpace(parser.aeroURL.GetURL().Query().Get("rack_aware"))

	if rackAwareStr != "" {
		rackAware, _ := strconv.ParseBool(rackAwareStr)
		parser.policy.RackAware = rackAware
	}
}

func (parser *ClientPolicyParser) RackId() {
	rackIdStr := strings.TrimSpace(parser.aeroURL.GetURL().Query().Get("rack_id"))

	if rackIdStr != "" {
		rackId, _ := strconv.Atoi(rackIdStr)
		parser.policy.RackId = rackId
	}
}

func (parser *ClientPolicyParser) IgnoreOtherSubnetAliases() {
	ignoreOtherSubnetAliasesStr := strings.TrimSpace(parser.aeroURL.GetURL().Query().Get("ignore_subnet_aliases"))

	if ignoreOtherSubnetAliasesStr != "" {
		ignoreOtherSubnetAliases, _ := strconv.ParseBool(ignoreOtherSubnetAliasesStr)
		parser.policy.IgnoreOtherSubnetAliases = ignoreOtherSubnetAliases
	}
}

func (parser *ClientPolicyParser) SeedOnlyCluster() {
	seedOnlyClusterStr := strings.TrimSpace(parser.aeroURL.GetURL().Query().Get("seed_only_cluster"))

	if seedOnlyClusterStr != "" {
		seedOnlyCluster, _ := strconv.ParseBool(seedOnlyClusterStr)
		parser.policy.SeedOnlyCluster = seedOnlyCluster
	}
}
