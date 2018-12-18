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

// PromMetrics contains metrics exposed by Consensus.
type Metrics struct {
	ApigatewayReceivedTx metrics.Counter
	SwitchTakenTx        metrics.Counter
	TxpoolIngressTx      metrics.Counter
	TxpoolPooledTx       metrics.Counter
	TxpoolDiscardedTx    metrics.Counter
	TxpoolDuplacatedTx   metrics.Counter
	TxpoolOutgoingTx     metrics.Counter
	BlockHeight          metrics.Gauge
	BlockTxNum           metrics.Gauge
	CommittedTx          metrics.Counter
}

func init() {
	// NopConsensusMetrics returns no-op Metrics.
	// Used by default.
	JTMetrics = &Metrics{
		ApigatewayReceivedTx: discard.NewCounter(),
		SwitchTakenTx:        discard.NewCounter(),
		TxpoolIngressTx:      discard.NewCounter(),
		TxpoolPooledTx:       discard.NewCounter(),
		TxpoolDiscardedTx:    discard.NewCounter(),
		TxpoolDuplacatedTx:   discard.NewCounter(),
		TxpoolOutgoingTx:     discard.NewCounter(),
		BlockHeight:          discard.NewGauge(),
		BlockTxNum:           discard.NewGauge(),
		CommittedTx:          discard.NewCounter(),
	}
}

// createJTMetrics creates Metrics build using Prometheus client library.
func createMetrics() {
	JTMetrics = &Metrics{
		ApigatewayReceivedTx: kitprometheus.NewCounterFrom(prometheus.CounterOpts{
			Subsystem: "apigateway",
			Name:      "received_tx",
			Help:      "Accumulated num of tx received by apigateway.",
		}, []string{}),
		SwitchTakenTx: kitprometheus.NewCounterFrom(prometheus.CounterOpts{
			Subsystem: "gossipswitch",
			Name:      "taken_tx",
			Help:      "Accumulated num of tx taken by gossipswitch.",
		}, []string{}),
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
	JTMetrics.ApigatewayReceivedTx.Add(0)
	JTMetrics.SwitchTakenTx.Add(0)
	JTMetrics.TxpoolIngressTx.Add(0)
	JTMetrics.TxpoolPooledTx.Add(0)
	JTMetrics.TxpoolDuplacatedTx.Add(0)
	JTMetrics.TxpoolDiscardedTx.Add(0)
	JTMetrics.TxpoolOutgoingTx.Add(0)
	JTMetrics.BlockHeight.Add(0)
	JTMetrics.BlockTxNum.Add(0)
	JTMetrics.CommittedTx.Add(0)
}
