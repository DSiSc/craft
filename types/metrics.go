/*
 * Define types and structures related to prometheus metrics.
 */
package types

import (
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/discard"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus"
)

// PromMetrics contains metrics exposed by Consensus.
type JTMetrics struct {
	Height  metrics.Gauge
	NumTx   metrics.Gauge
	TotalTx metrics.Gauge
}

// PromConsensusMetrics returns consensus Metrics build using Prometheus client library.
func PromJTMetrics() *JTMetrics {
	return &JTMetrics{
		Height: kitprometheus.NewGaugeFrom(prometheus.GaugeOpts{
			Subsystem: "consensus",
			Name:      "height",
			Help:      "Height of blocks.",
		}, []string{}),
		NumTx: kitprometheus.NewGaugeFrom(prometheus.GaugeOpts{
			Subsystem: "consensus",
			Name:      "num_tx",
			Help:      "Num of tx in current block.",
		}, []string{}),
		TotalTx: kitprometheus.NewGaugeFrom(prometheus.GaugeOpts{
			Subsystem: "consensus",
			Name:      "total_tx",
			Help:      "Total num of tx.",
		}, []string{}),
	}
}

// NopConsensusMetrics returns no-op Metrics.
func NopJTMetrics() *JTMetrics {
	return &JTMetrics{
		Height:  discard.NewGauge(),
		NumTx:   discard.NewGauge(),
		TotalTx: discard.NewGauge(),
	}
}

