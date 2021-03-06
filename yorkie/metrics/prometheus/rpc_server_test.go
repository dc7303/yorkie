package prometheus

import (
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus/testutil"
)

func TestRPCServerMetrics(t *testing.T) {
	var rpcServerMetrics = NewRPCServerMetrics()

	t.Run("observe pushpull response seconds test", func(t *testing.T) {
		rpcServerMetrics.ObservePushpullResponseSeconds(3)
		rpcServerMetrics.ObservePushpullResponseSeconds(5)

		expected := `
			# HELP yorkie_rpcserver_pushpull_response_seconds Response time of PushPull API.
            # TYPE yorkie_rpcserver_pushpull_response_seconds histogram
			yorkie_rpcserver_pushpull_response_seconds_bucket{le="0.005"} 0
            yorkie_rpcserver_pushpull_response_seconds_bucket{le="0.01"} 0
            yorkie_rpcserver_pushpull_response_seconds_bucket{le="0.025"} 0
            yorkie_rpcserver_pushpull_response_seconds_bucket{le="0.05"} 0
            yorkie_rpcserver_pushpull_response_seconds_bucket{le="0.1"} 0
            yorkie_rpcserver_pushpull_response_seconds_bucket{le="0.25"} 0
            yorkie_rpcserver_pushpull_response_seconds_bucket{le="0.5"} 0
            yorkie_rpcserver_pushpull_response_seconds_bucket{le="1"} 0
            yorkie_rpcserver_pushpull_response_seconds_bucket{le="2.5"} 0
            yorkie_rpcserver_pushpull_response_seconds_bucket{le="5"} 2
            yorkie_rpcserver_pushpull_response_seconds_bucket{le="10"} 2
            yorkie_rpcserver_pushpull_response_seconds_bucket{le="+Inf"} 2
            yorkie_rpcserver_pushpull_response_seconds_sum 8
            yorkie_rpcserver_pushpull_response_seconds_count 2
		`
		if err := testutil.CollectAndCompare(rpcServerMetrics.pushpullResponseSeconds, strings.NewReader(expected)); err != nil {
			t.Errorf("unexpected collecting result:\n%s", err)
		}
	})
}
