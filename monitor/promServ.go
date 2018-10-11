package monitor

import (
	"context"
	"github.com/DSiSc/craft/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var promServ *http.Server

// startPrometheusServer starts a Prometheus HTTP server, listening for metrics
// collectors on addr.
func StartPrometheusServer(port string, maxConn int) {

	createMetrics()

	// create prometheus server
	if promServ == nil {
		promServ = &http.Server{
			//Addr: ":" + node.config.PrometheusPort,
			Addr: ":" + port,
			Handler: promhttp.InstrumentMetricHandler(
				prometheus.DefaultRegisterer, promhttp.HandlerFor(
					prometheus.DefaultGatherer,
					promhttp.HandlerOpts{MaxRequestsInFlight: maxConn},
				),
			),
		}
	}

	// start prometheus server
	go func() {
		if err := promServ.ListenAndServe(); err != http.ErrServerClosed {
			// Error starting or closing listener:
			log.Error("Prometheus HTTP server ListenAndServe", "err", err)
		}
	}()
}

// stopPrometheusServer stops a Prometheus HTTP server
func StopPrometheusServer() {
	if promServ != nil {
		if err := promServ.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.ErrorKV("Prometheus HTTP server Shutdown", map[string]interface{}{"err": err})
		}
	}
}
