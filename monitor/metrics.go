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
		TxpoolIngressTx:    discard.NewCounter(),
		TxpoolPooledTx:     discard.NewCounter(),
		TxpoolDiscardedTx:  discard.NewCounter(),
		TxpoolDuplacatedTx: discard.NewCounter(),
		TxpoolOutgoingTx:   discard.NewCounter(),
		BlockHeight:        discard.NewGauge(),
		BlockTxNum:         discard.NewGauge(),
		CommittedTx:        discard.NewCounter(),
	}
}

// PromMetrics contains metrics exposed by Consensus.
type Metrics struct {
	TxpoolIngressTx    metrics.Counter
	TxpoolPooledTx     metrics.Counter
	TxpoolDiscardedTx  metrics.Counter
	TxpoolDuplacatedTx metrics.Counter
	TxpoolOutgoingTx   metrics.Counter
	BlockHeight        metrics.Gauge
	BlockTxNum         metrics.Gauge
	CommittedTx        metrics.Counter
}

// createJTMetrics creates Metrics build using Prometheus client library.
func createMetrics() {
	JTMetrics = &Metrics{
		TxpoolIngressTx: kitprometheus.NewCounterFrom(prometheus.CounterOpts{
			Subsystem: "txpool",
			Name:      "ingress_tx",
			Help:      "Accumulated num of incoming tx to txpool.",
		}, []string{}),
		TxpoolPooledTx: kitprometheus.NewCounterFrom(prometheus.CounterOpts{
			Subsystem: "txpool",
			Name:      "pooled_tx",
			Help:      "Accumulated num of tx pooled into txpool.",
		}, []string{}),
		TxpoolDiscardedTx: kitprometheus.NewCounterFrom(prometheus.CounterOpts{
			Subsystem: "txpool",
			Name:      "discarded_tx",
			Help:      "Accumulated num of discarded tx because txpool is full.",
		}, []string{}),
		TxpoolDuplacatedTx: kitprometheus.NewCounterFrom(prometheus.CounterOpts{
			Subsystem: "txpool",
			Name:      "duplicated_tx",
			Help:      "Accumulated num of duplicated tx.",
		}, []string{}),
		TxpoolOutgoingTx: kitprometheus.NewCounterFrom(prometheus.CounterOpts{
			Subsystem: "txpool",
			Name:      "outgoing_tx",
			Help:      "Accumulated num of tx out from txpool.",
		}, []string{}),
		BlockHeight: kitprometheus.NewGaugeFrom(prometheus.GaugeOpts{
			Subsystem: "store",
			Name:      "height",
			Help:      "Height of blocks.",
		}, []string{}),
		BlockTxNum: kitprometheus.NewGaugeFrom(prometheus.GaugeOpts{
			Subsystem: "store",
			Name:      "block_tx_num",
			Help:      "Num of tx contained in latest block.",
		}, []string{}),
		CommittedTx: kitprometheus.NewCounterFrom(prometheus.CounterOpts{
			Subsystem: "store",
			Name:      "total_tx_num",
			Help:      "Accumulated num of committed tx.",
		}, []string{}),
	}
}
