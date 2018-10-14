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
		MempoolIngressTx:  discard.NewCounter(),
		MempoolPooledTx:   discard.NewCounter(),
		MempoolOutgoingTx: discard.NewCounter(),
		BlockHeight:       discard.NewGauge(),
		BlockTxNum:        discard.NewGauge(),
		CommittedTx:       discard.NewCounter(),
	}
}

// PromMetrics contains metrics exposed by Consensus.
type Metrics struct {
	MempoolIngressTx  metrics.Counter
	MempoolPooledTx   metrics.Counter
	MempoolOutgoingTx metrics.Counter
	BlockHeight       metrics.Gauge
	BlockTxNum        metrics.Gauge
	CommittedTx       metrics.Counter
}

// createJTMetrics creates Metrics build using Prometheus client library.
func createMetrics() {
	JTMetrics = &Metrics{
		MempoolIngressTx: kitprometheus.NewCounterFrom(prometheus.CounterOpts{
			Subsystem: "mempool",
			Name:      "ingress_tx",
			Help:      "Accumulated num of incoming tx mempool.",
		}, []string{}),
		MempoolPooledTx: kitprometheus.NewCounterFrom(prometheus.CounterOpts{
			Subsystem: "mempool",
			Name:      "pooled_tx",
			Help:      "Accumulated num of tx pooled into mempool.",
		}, []string{}),
		MempoolOutgoingTx: kitprometheus.NewCounterFrom(prometheus.CounterOpts{
			Subsystem: "mempool",
			Name:      "outgoing_tx",
			Help:      "Accumulated num of tx out from mempool.",
		}, []string{}),
		BlockHeight: kitprometheus.NewGaugeFrom(prometheus.GaugeOpts{
			Subsystem: "store",
			Name:      "height",
			Help:      "Height of blocks.",
		}, []string{}),
		BlockTxNum: kitprometheus.NewGaugeFrom(prometheus.GaugeOpts{
			Subsystem: "store",
			Name:      "block_tx_num",
			Help:      "Num of tx contained in recent block.",
		}, []string{}),
		CommittedTx: kitprometheus.NewCounterFrom(prometheus.CounterOpts{
			Subsystem: "store",
			Name:      "total_tx_num",
			Help:      "Accumulated num of committed tx.",
		}, []string{}),
	}
}
