package metrics

import "github.com/yorkie-team/yorkie/yorkie/metrics/prometheus"

type metrics struct {
	Server    ServerMetrics
	RPCServer RPCServerMetrics
}

func newMetrics() *metrics {
	return &metrics{
		Server:    prometheus.NewServerMetrics(),
		RPCServer: prometheus.NewRPCServerMetrics(),
	}
}

type ServerMetrics interface {
	WithServerVersion(version string)
}

type RPCServerMetrics interface {
	ObservePushpullResponseSeconds(seconds float64)
	AddPushpullReceivedChanges(count float64)
	AddPushpullSentChanges(count float64)
	ObservePushpullSnapshotDurationSeconds(seconds float64)
	AddPushpullSnapshotBytes(byte float64)
}