package monitor

import (
	"context"
	"github.com/DSiSc/craft/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var prometheusServ *http.Server

type PrometheusConfig struct {
	PrometheusEnabled bool
	PrometheusPort    string
	PrometheusMaxConn int
}

// startPrometheusServer starts a Prometheus HTTP server, listening for metrics
// collectors on addr.
func StartPrometheusServer(config PrometheusConfig) {

	if !config.PrometheusEnabled {
		return
	}

	createMetrics()

	// create prometheus server
	if prometheusServ == nil {
		prometheusServ = &http.Server{
			Addr: ":" + config.PrometheusPort,
			Handler: promhttp.InstrumentMetricHandler(
				prometheus.DefaultRegisterer, promhttp.HandlerFor(
					prometheus.DefaultGatherer,
					promhttp.HandlerOpts{MaxRequestsInFlight: config.PrometheusMaxConn},
				),
			),
		}
	}

	// start prometheus server
	go func() {
		if err := prometheusServ.ListenAndServe(); err != http.ErrServerClosed {
			// Error starting or closing listener:
			log.Error("Prometheus HTTP server ListenAndServe", "err", err)
		}
	}()
}

// stopPrometheusServer stops a Prometheus HTTP server
func StopPrometheusServer() {
	if prometheusServ != nil {
		if err := prometheusServ.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.ErrorKV("Prometheus HTTP server Shutdown", map[string]interface{}{"err": err})
		}
	}
}
