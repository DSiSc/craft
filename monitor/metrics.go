/*
 * Define types and structures related to prometheus metrics.
 */
package monitor

import (
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/discard"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus"
)

var JTMetrics *Metrics

func init() {
	// NopConsensusMetrics returns no-op Metrics.
	// Used by default.
	JTMetrics = &Metrics{
		Height:  discard.NewGauge(),
		NumTx:   discard.NewGauge(),
		TotalTx: discard.NewCounter(),
	}
}

// PromMetrics contains metrics exposed by Consensus.
type Metrics struct {
	Height  metrics.Gauge
	NumTx   metrics.Gauge
	TotalTx metrics.Counter
}

// createJTMetrics creates Metrics build using Prometheus client library.
func createMetrics() {
	JTMetrics = &Metrics{
		Height: kitprometheus.NewGaugeFrom(prometheus.GaugeOpts{
			Subsystem: "store",
			Name:      "height",
			Help:      "Height of blocks.",
		}, []string{}),
		NumTx: kitprometheus.NewGaugeFrom(prometheus.GaugeOpts{
			Subsystem: "store",
			Name:      "block_tx_num",
			Help:      "Num of tx in current block.",
		}, []string{}),
		TotalTx: kitprometheus.NewCounterFrom(prometheus.CounterOpts{
			Subsystem: "store",
			Name:      "total_tx_num",
			Help:      "Total num of tx.",
		}, []string{}),
	}
}
