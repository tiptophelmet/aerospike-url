package clientpolicy

import (
	"fmt"
	"testing"
	"time"

	"github.com/aerospike/aerospike-client-go/v6"
	"github.com/tiptophelmet/aerospike-url/aerofactory"
	"github.com/tiptophelmet/aerospike-url/aerourl"
)

func TestClientPolicyParser_AuthModeInternal(t *testing.T) {
	aeroURL, _ := aerourl.Init("aerospike://127.0.0.1:3000/aero-namespace-001?auth_mode=auth_mode_internal")
	policy := aerospike.NewClientPolicy()

	parser := &ClientPolicyParser{aeroURL, policy}
	parser.AuthMode()

	if parser.policy.AuthMode != aerospike.AuthModeInternal {
		t.Fatalf("got: %v, want: parser.policy.AuthModeInternal = 0 (aerospike.AuthModeInternal)", parser.policy.AuthMode)
	}
}

func TestClientPolicyParser_AuthModeExternal(t *testing.T) {
	aeroURL, _ := aerourl.Init("aerospike://127.0.0.1:3000/aero-namespace-001?auth_mode=auth_mode_external")
	policy := aerospike.NewClientPolicy()

	parser := &ClientPolicyParser{aeroURL, policy}
	parser.AuthMode()

	if parser.policy.AuthMode != aerospike.AuthModeExternal {
		t.Fatalf("got: %v, want: parser.policy.AuthMode = 1 (aerospike.AuthModeExternal)", parser.policy.AuthMode)
	}
}

func TestClientPolicyParser_AuthModePKI(t *testing.T) {
	aeroURL, _ := aerourl.Init("aerospike://127.0.0.1:3000/aero-namespace-001?auth_mode=auth_mode_pki")
	policy := aerospike.NewClientPolicy()

	parser := &ClientPolicyParser{aeroURL, policy}
	parser.AuthMode()

	if parser.policy.AuthMode != aerospike.AuthModePKI {
		t.Fatalf("got: %v, want: parser.policy.AuthMode = 2 (aerospike.AuthModePKI)", parser.policy.AuthMode)
	}
}

func TestClientPolicyParser_User(t *testing.T) {
	user := "aero-user-001"
	connStr := fmt.Sprintf("aerospike://%v:aerouserpassw123@127.0.0.1:3000/aero-namespace-001", user)

	aeroURL, _ := aerourl.Init(connStr)
	policy := aerospike.NewClientPolicy()

	parser := &ClientPolicyParser{aeroURL, policy}
	parser.User()

	if parser.policy.User != user {
		t.Fatalf("got: %s, want: %s", parser.policy.User, user)
	}
}

func TestClientPolicyParser_Password(t *testing.T) {
	password := "aerouserpassw123"
	connStr := fmt.Sprintf("aerospike://aero-user-001:%v@127.0.0.1:3000/aero-namespace-001", password)

	aeroURL, _ := aerourl.Init(connStr)
	policy := aerospike.NewClientPolicy()

	parser := &ClientPolicyParser{aeroURL, policy}
	parser.Password()

	if parser.policy.Password != password {
		t.Fatalf("got: %s, want: %s", parser.policy.Password, password)
	}
}

func TestClientPolicyParser_ClusterName(t *testing.T) {
	clusterName := "aero-cluster-001"
	connStr := fmt.Sprintf("aerospike://127.0.0.1:3000/aero-namespace-001?cluster_name=%v", clusterName)

	aeroURL, _ := aerourl.Init(connStr)
	policy := aerospike.NewClientPolicy()

	parser := &ClientPolicyParser{aeroURL, policy}
	parser.ClusterName()

	if parser.policy.ClusterName != clusterName {
		t.Fatalf("got: %s, want: %s", parser.policy.ClusterName, clusterName)
	}
}

func TestClientPolicyParser_Timeout(t *testing.T) {
	timeoutStr := "60s"
	connStr := fmt.Sprintf("aerospike://127.0.0.1:3000/aero-namespace-001?timeout=%v", timeoutStr)

	aeroURL, _ := aerourl.Init(connStr)
	policy := aerospike.NewClientPolicy()

	parser := &ClientPolicyParser{aeroURL, policy}
	parser.Timeout()

	timeout, _ := time.ParseDuration(timeoutStr)
	if parser.policy.Timeout != timeout {
		t.Fatalf("got: %v, want: %v", parser.policy.Timeout, timeout)
	}
}

func TestClientPolicyParser_IdleTimeout(t *testing.T) {
	idleTimeoutStr := "15s"
	connStr := fmt.Sprintf("aerospike://127.0.0.1:3000/aero-namespace-001?idle_timeout=%v", idleTimeoutStr)

	aeroURL, _ := aerourl.Init(connStr)
	policy := aerospike.NewClientPolicy()

	parser := &ClientPolicyParser{aeroURL, policy}
	parser.IdleTimeout()

	idleTimeout, _ := time.ParseDuration(idleTimeoutStr)
	if parser.policy.IdleTimeout != idleTimeout {
		t.Fatalf("got: %v, want: %v", parser.policy.IdleTimeout, idleTimeout)
	}
}

func TestClientPolicyParser_LoginTimeout(t *testing.T) {
	loginTimeoutStr := "10s"
	connStr := fmt.Sprintf("aerospike://127.0.0.1:3000/aero-namespace-001?login_timeout=%v", loginTimeoutStr)

	aeroURL, _ := aerourl.Init(connStr)
	policy := aerospike.NewClientPolicy()

	parser := &ClientPolicyParser{aeroURL, policy}
	parser.LoginTimeout()

	loginTimeout, _ := time.ParseDuration(loginTimeoutStr)
	if parser.policy.LoginTimeout != loginTimeout {
		t.Fatalf("got: %v, want: %v", parser.policy.LoginTimeout, loginTimeout)
	}
}

func TestClientPolicyParser_ConnectionQueueSize(t *testing.T) {
	connQueueSize := 100
	connStr := fmt.Sprintf("aerospike://127.0.0.1:3000/aero-namespace-001?connection_queue_size=%d", connQueueSize)

	aeroURL, _ := aerourl.Init(connStr)
	policy := aerospike.NewClientPolicy()

	parser := &ClientPolicyParser{aeroURL, policy}
	parser.ConnectionQueueSize()

	if parser.policy.ConnectionQueueSize != connQueueSize {
		t.Fatalf("got: %v, want: %v", parser.policy.ConnectionQueueSize, connQueueSize)
	}
}

func TestClientPolicyParser_MinConnectionsPerNode(t *testing.T) {
	minConnsPerNode := 5
	connStr := fmt.Sprintf("aerospike://127.0.0.1:3000/aero-namespace-001?min_connections_per_node=%d", minConnsPerNode)

	aeroURL, _ := aerourl.Init(connStr)
	policy := aerospike.NewClientPolicy()

	parser := &ClientPolicyParser{aeroURL, policy}
	parser.MinConnectionsPerNode()

	if parser.policy.MinConnectionsPerNode != minConnsPerNode {
		t.Fatalf("got: %v, want: %v", parser.policy.MinConnectionsPerNode, minConnsPerNode)
	}
}

func TestClientPolicyParser_MaxErrorRate(t *testing.T) {
	maxErrorRate := 10
	connStr := fmt.Sprintf("aerospike://127.0.0.1:3000/aero-namespace-001?max_error_rate=%d", maxErrorRate)

	aeroURL, _ := aerourl.Init(connStr)
	policy := aerospike.NewClientPolicy()

	parser := &ClientPolicyParser{aeroURL, policy}
	parser.MaxErrorRate()

	if parser.policy.MaxErrorRate != maxErrorRate {
		t.Fatalf("got: %v, want: %v", parser.policy.MaxErrorRate, maxErrorRate)
	}
}

func TestClientPolicyParser_ErrorRateWindow(t *testing.T) {
	errorRateWindow := 100
	connStr := fmt.Sprintf("aerospike://127.0.0.1:3000/aero-namespace-001?error_rate_window=%d", errorRateWindow)

	aeroURL, _ := aerourl.Init(connStr)
	policy := aerospike.NewClientPolicy()

	parser := &ClientPolicyParser{aeroURL, policy}
	parser.ErrorRateWindow()

	if parser.policy.ErrorRateWindow != errorRateWindow {
		t.Fatalf("got: %v, want: %v", parser.policy.ErrorRateWindow, errorRateWindow)
	}
}

func TestClientPolicyParser_LimitConnectionsToQueueSize(t *testing.T) {
	limitConnsToQueueSize := true
	connStr := fmt.Sprintf("aerospike://127.0.0.1:3000/aero-namespace-001?limit_connections_to_queue_size=%v", limitConnsToQueueSize)

	aeroURL, _ := aerourl.Init(connStr)
	policy := aerospike.NewClientPolicy()

	parser := &ClientPolicyParser{aeroURL, policy}
	parser.LimitConnectionsToQueueSize()

	if parser.policy.LimitConnectionsToQueueSize != limitConnsToQueueSize {
		t.Fatalf("got: %v, want: %v", parser.policy.LimitConnectionsToQueueSize, limitConnsToQueueSize)
	}
}

func TestClientPolicyParser_OpeningConnectionThreshold(t *testing.T) {
	openingConnThreshold := 50
	connStr := fmt.Sprintf("aerospike://127.0.0.1:3000/aero-namespace-001?opening_connection_threshold=%d", openingConnThreshold)

	aeroURL, _ := aerourl.Init(connStr)
	policy := aerospike.NewClientPolicy()

	parser := &ClientPolicyParser{aeroURL, policy}
	parser.OpeningConnectionThreshold()

	if parser.policy.OpeningConnectionThreshold != openingConnThreshold {
		t.Fatalf("got: %v, want: %v", parser.policy.OpeningConnectionThreshold, openingConnThreshold)
	}
}

func TestClientPolicyParser_FailIfNotConnected(t *testing.T) {
	failIfNotConnected := true
	connStr := fmt.Sprintf("aerospike://127.0.0.1:3000/aero-namespace-001?fail_if_not_connected=%v", failIfNotConnected)

	aeroURL, _ := aerourl.Init(connStr)
	policy := aerospike.NewClientPolicy()

	parser := &ClientPolicyParser{aeroURL, policy}
	parser.FailIfNotConnected()

	if parser.policy.FailIfNotConnected != failIfNotConnected {
		t.Fatalf("got: %v, want: %v", parser.policy.FailIfNotConnected, failIfNotConnected)
	}
}

func TestClientPolicyParser_TendInterval(t *testing.T) {
	tendIntervalStr := "5s"
	connStr := fmt.Sprintf("aerospike://127.0.0.1:3000/aero-namespace-001?tend_interval=%v", tendIntervalStr)

	aeroURL, _ := aerourl.Init(connStr)
	policy := aerospike.NewClientPolicy()

	parser := &ClientPolicyParser{aeroURL, policy}
	parser.TendInterval()

	tendInterval, _ := time.ParseDuration(tendIntervalStr)
	if parser.policy.TendInterval != tendInterval {
		t.Fatalf("got: %v, want: %v", parser.policy.TendInterval, tendInterval)
	}
}

func TestClientPolicyParser_UseServicesAlternate(t *testing.T) {
	useServicesAlternate := true
	connStr := fmt.Sprintf("aerospike://127.0.0.1:3000/aero-namespace-001?use_services_alternate=%v", useServicesAlternate)

	aeroURL, _ := aerourl.Init(connStr)
	policy := aerospike.NewClientPolicy()

	parser := &ClientPolicyParser{aeroURL, policy}
	parser.UseServicesAlternate()

	if parser.policy.UseServicesAlternate != useServicesAlternate {
		t.Fatalf("got: %v, want: %v", parser.policy.UseServicesAlternate, useServicesAlternate)
	}
}

func TestClientPolicyParser_RackAware(t *testing.T) {
	rackAware := true
	connStr := fmt.Sprintf("aerospike://127.0.0.1:3000/aero-namespace-001?rack_aware=%v", rackAware)

	aeroURL, _ := aerourl.Init(connStr)
	policy := aerospike.NewClientPolicy()

	parser := &ClientPolicyParser{aeroURL, policy}
	parser.RackAware()

	if parser.policy.RackAware != rackAware {
		t.Fatalf("got: %v, want: %v", parser.policy.RackAware, rackAware)
	}
}

func TestClientPolicyParser_RackId(t *testing.T) {
	rackId := 2
	connStr := fmt.Sprintf("aerospike://127.0.0.1:3000/aero-namespace-001?rack_id=%d", rackId)

	aeroURL, _ := aerourl.Init(connStr)
	policy := aerospike.NewClientPolicy()

	parser := &ClientPolicyParser{aeroURL, policy}
	parser.RackId()

	if parser.policy.RackId != rackId {
		t.Fatalf("got: %v, want: %v", parser.policy.RackId, rackId)
	}
}

func TestClientPolicyParser_IgnoreOtherSubnetAliases(t *testing.T) {
	ignoreOtherSubnetAliases := true
	connStr := fmt.Sprintf("aerospike://127.0.0.1:3000/aero-namespace-001?ignore_subnet_aliases=%v", ignoreOtherSubnetAliases)

	aeroURL, _ := aerourl.Init(connStr)
	policy := aerospike.NewClientPolicy()

	parser := &ClientPolicyParser{aeroURL, policy}
	parser.IgnoreOtherSubnetAliases()

	if parser.policy.IgnoreOtherSubnetAliases != ignoreOtherSubnetAliases {
		t.Fatalf("got: %v, want: %v", parser.policy.IgnoreOtherSubnetAliases, ignoreOtherSubnetAliases)
	}
}

func TestClientPolicyParser_SeedOnlyCluster(t *testing.T) {
	seedOnlyCluster := true
	connStr := fmt.Sprintf("aerospike://127.0.0.1:3000/aero-namespace-001?seed_only_cluster=%v", seedOnlyCluster)

	aeroURL, _ := aerourl.Init(connStr)
	policy := aerospike.NewClientPolicy()

	parser := &ClientPolicyParser{aeroURL, policy}
	parser.SeedOnlyCluster()

	if parser.policy.SeedOnlyCluster != seedOnlyCluster {
		t.Fatalf("got: %v, want: %v", parser.policy.SeedOnlyCluster, seedOnlyCluster)
	}
}

func TestClientPolicyParser_GetClientPolicy(t *testing.T) {
	aeroURL, _ := aerourl.Init("aerospike://127.0.0.1:3000/aero-namespace-001?auth_mode=auth_mode_pki")
	policy := aerospike.NewClientPolicy()

	parser := &ClientPolicyParser{aeroURL, policy}

	result := parser.GetClientPolicy()

	if result != policy {
		t.Fatalf("got: %v, want: %v", result, policy)
	}
}

func TestParse(t *testing.T) {
	var (
		user                  = "aero-user-001"
		password              = "aerouserpassw123"
		authMode              = "auth_mode_pki"
		timeout               = "5s"
		clusterName           = "myCluster"
		minConnectionsPerNode = 3
	)

	connStr := fmt.Sprintf("aerospike://%s:%s@127.0.0.1:3000/aero-namespace-001?auth_mode=%s&timeout=%s&cluster_name=%s&min_connections_per_node=%d",
		user, password, authMode, timeout, clusterName, minConnectionsPerNode)

	aeroURL, _ := aerourl.Init(connStr)

	clientFactory := &aerofactory.AerospikeClientFactory{}
	clientFactory.SetHostname("127.0.0.1")
	clientFactory.SetPort(3000)

	Parse(aeroURL, clientFactory)

	if clientFactory.GetClientPolicy().AuthMode != aerospike.AuthModePKI {
		t.Fatalf("got: %d, want: clientFactory.GetClientPolicy().AuthMode = 2 (aerospike.AuthModePKI)", clientFactory.GetClientPolicy().AuthMode)
	}

	if clientFactory.GetClientPolicy().User != user {
		t.Fatalf("got: %s, want: %s", clientFactory.GetClientPolicy().User, user)
	}

	if clientFactory.GetClientPolicy().Password != password {
		t.Fatalf("got: %s, want: %s", clientFactory.GetClientPolicy().Password, password)
	}

	if clientFactory.GetClientPolicy().Timeout.String() != timeout {
		t.Fatalf("got: %s, want: %s", clientFactory.GetClientPolicy().Timeout.String(), timeout)
	}

	if clientFactory.GetClientPolicy().ClusterName != clusterName {
		t.Fatalf("got: %s, want: %s", clientFactory.GetClientPolicy().ClusterName, clusterName)
	}

	if clientFactory.GetClientPolicy().MinConnectionsPerNode != minConnectionsPerNode {
		t.Fatalf("got: %d, want: %d", clientFactory.GetClientPolicy().MinConnectionsPerNode, minConnectionsPerNode)
	}
}
