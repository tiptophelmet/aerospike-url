package aerospikeurl

import (
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/aerospike/aerospike-client-go/v6"
)

func parseClientPolicyQuery(connURL *url.URL, client *AerospikeClientFactory) {
	parseAuthMode(connURL, client.policy)
	parseUser(connURL, client.policy)
	parsePassword(connURL, client.policy)
	parseClusterName(connURL, client.policy)
	parseTimeout(connURL, client.policy)
	parseIdleTimeout(connURL, client.policy)
	parseLoginTimeout(connURL, client.policy)
	parseConnectionQueueSize(connURL, client.policy)
	parseMinConnectionsPerNode(connURL, client.policy)
	parseMaxErrorRate(connURL, client.policy)
	parseErrorRateWindow(connURL, client.policy)
	parseLimitConnectionsToQueueSize(connURL, client.policy)
	parseOpeningConnectionThreshold(connURL, client.policy)
	parseFailIfNotConnected(connURL, client.policy)
	parseTendInterval(connURL, client.policy)
	parseUseServicesAlternate(connURL, client.policy)
	parseRackAware(connURL, client.policy)
	parseRackId(connURL, client.policy)
	parseIgnoreOtherSubnetAliases(connURL, client.policy)
	parseSeedOnlyCluster(connURL, client.policy)
}

func parseAuthMode(connURL *url.URL, policy *aerospike.ClientPolicy) {
	authMode := connURL.Query().Get("auth_mode")

	if authMode == "auth_mode_internal" {
		policy.AuthMode = aerospike.AuthModeInternal
	} else if authMode == "auth_mode_external" {
		policy.AuthMode = aerospike.AuthModeExternal
	} else if authMode == "auth_mode_pki" {
		policy.AuthMode = aerospike.AuthModePKI
	}
}

func parseUser(connURL *url.URL, policy *aerospike.ClientPolicy) {
	user := connURL.User.Username()
	if user != "" {
		policy.User = user
	}
}

func parsePassword(connURL *url.URL, policy *aerospike.ClientPolicy) {
	password, isPasswordSet := connURL.User.Password()
	if isPasswordSet {
		policy.Password = password
	}
}

func parseClusterName(connURL *url.URL, policy *aerospike.ClientPolicy) {
	clusterName := connURL.Query().Get("cluster_name")
	if clusterName != "" {
		policy.ClusterName = clusterName
	}
}

func parseTimeout(connURL *url.URL, policy *aerospike.ClientPolicy) {
	timeoutStr := connURL.Query().Get("timeout")

	if timeoutStr != "" {
		timeout, _ := time.ParseDuration(timeoutStr)
		policy.Timeout = timeout
	}
}

func parseIdleTimeout(connURL *url.URL, policy *aerospike.ClientPolicy) {
	idleTimeoutStr := connURL.Query().Get("idle_timeout")

	if idleTimeoutStr != "" {
		idleTimeout, _ := time.ParseDuration(idleTimeoutStr)
		policy.IdleTimeout = idleTimeout
	}
}

func parseLoginTimeout(connURL *url.URL, policy *aerospike.ClientPolicy) {
	loginTimeoutStr := connURL.Query().Get("login_timeout")

	if loginTimeoutStr != "" {
		loginTimeout, _ := time.ParseDuration(loginTimeoutStr)
		policy.LoginTimeout = loginTimeout
	}
}

func parseConnectionQueueSize(connURL *url.URL, policy *aerospike.ClientPolicy) {
	connQueueSizeStr := strings.TrimSpace(connURL.Query().Get("connection_queue_size"))

	if connQueueSizeStr != "" {
		connQueueSize, _ := strconv.Atoi(connQueueSizeStr)
		policy.ConnectionQueueSize = connQueueSize
	}
}

func parseMinConnectionsPerNode(connURL *url.URL, policy *aerospike.ClientPolicy) {
	minConnsPerNodeStr := strings.TrimSpace(connURL.Query().Get("min_connections_per_node"))

	if minConnsPerNodeStr != "" {
		minConnsPerNode, _ := strconv.Atoi(minConnsPerNodeStr)
		policy.MinConnectionsPerNode = minConnsPerNode
	}
}

func parseMaxErrorRate(connURL *url.URL, policy *aerospike.ClientPolicy) {
	maxErrorRateStr := strings.TrimSpace(connURL.Query().Get("max_error_rate"))

	if maxErrorRateStr != "" {
		maxErrorRate, _ := strconv.Atoi(maxErrorRateStr)
		policy.MaxErrorRate = maxErrorRate
	}
}

func parseErrorRateWindow(connURL *url.URL, policy *aerospike.ClientPolicy) {
	errorRateWindowStr := strings.TrimSpace(connURL.Query().Get("error_rate_window"))

	if errorRateWindowStr != "" {
		errorRateWindow, _ := strconv.Atoi(errorRateWindowStr)
		policy.ErrorRateWindow = errorRateWindow
	}
}

func parseLimitConnectionsToQueueSize(connURL *url.URL, policy *aerospike.ClientPolicy) {
	limitConnsToQueueSizeStr := strings.TrimSpace(connURL.Query().Get("limit_connections_to_queue_size"))

	if limitConnsToQueueSizeStr != "" {
		limitConnsToQueueSize, _ := strconv.ParseBool(limitConnsToQueueSizeStr)
		policy.LimitConnectionsToQueueSize = limitConnsToQueueSize
	}
}

func parseOpeningConnectionThreshold(connURL *url.URL, policy *aerospike.ClientPolicy) {
	openingConnThresholdStr := strings.TrimSpace(connURL.Query().Get("opening_connection_threshold"))

	if openingConnThresholdStr != "" {
		openingConnThreshold, _ := strconv.Atoi(openingConnThresholdStr)
		policy.OpeningConnectionThreshold = openingConnThreshold
	}
}

func parseFailIfNotConnected(connURL *url.URL, policy *aerospike.ClientPolicy) {
	failIfNotConnectedStr := strings.TrimSpace(connURL.Query().Get("fail_if_not_connected"))

	if failIfNotConnectedStr != "" {
		failIfNotConnected, _ := strconv.ParseBool(failIfNotConnectedStr)
		policy.FailIfNotConnected = failIfNotConnected
	}
}

func parseTendInterval(connURL *url.URL, policy *aerospike.ClientPolicy) {
	tendIntervalStr := connURL.Query().Get("tend_interval")

	if tendIntervalStr != "" {
		tendInterval, _ := time.ParseDuration(tendIntervalStr)
		policy.TendInterval = tendInterval
	}
}

func parseUseServicesAlternate(connURL *url.URL, policy *aerospike.ClientPolicy) {
	useServicesAlternateStr := strings.TrimSpace(connURL.Query().Get("use_services_alternate"))

	if useServicesAlternateStr != "" {
		useServicesAlternate, _ := strconv.ParseBool(useServicesAlternateStr)
		policy.UseServicesAlternate = useServicesAlternate
	}
}

func parseRackAware(connURL *url.URL, policy *aerospike.ClientPolicy) {
	rackAwareStr := strings.TrimSpace(connURL.Query().Get("rack_aware"))

	if rackAwareStr != "" {
		rackAware, _ := strconv.ParseBool(rackAwareStr)
		policy.RackAware = rackAware
	}
}

func parseRackId(connURL *url.URL, policy *aerospike.ClientPolicy) {
	rackIdStr := strings.TrimSpace(connURL.Query().Get("rack_id"))

	if rackIdStr != "" {
		rackId, _ := strconv.Atoi(rackIdStr)
		policy.RackId = rackId
	}
}

func parseIgnoreOtherSubnetAliases(connURL *url.URL, policy *aerospike.ClientPolicy) {
	ignoreOtherSubnetAliasesStr := strings.TrimSpace(connURL.Query().Get("ignore_subnet_aliases"))

	if ignoreOtherSubnetAliasesStr != "" {
		ignoreOtherSubnetAliases, _ := strconv.ParseBool(ignoreOtherSubnetAliasesStr)
		policy.IgnoreOtherSubnetAliases = ignoreOtherSubnetAliases
	}
}

func parseSeedOnlyCluster(connURL *url.URL, policy *aerospike.ClientPolicy) {
	seedOnlyClusterStr := strings.TrimSpace(connURL.Query().Get("seed_only_cluster"))

	if seedOnlyClusterStr != "" {
		seedOnlyCluster, _ := strconv.ParseBool(seedOnlyClusterStr)
		policy.SeedOnlyCluster = seedOnlyCluster
	}
}
