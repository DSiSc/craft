package monitor

import (
	"expvar"
	"fmt"
	"net/http"
)

type ExpvarConfig struct {
	ExpvarEnabled bool
	ExpvarPort    string
	ExpvarPath    string
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	first := true
	report := func(key string, value interface{}) {
		if !first {
			fmt.Fprintf(w, ",\n")
		}
		first = false
		if str, ok := value.(string); ok {
			fmt.Fprintf(w, "%q: %q", key, str)
		} else {
			fmt.Fprintf(w, "%q: %v", key, value)
		}
	}

	fmt.Fprintf(w, "{\n")
	expvar.Do(func(kv expvar.KeyValue) {
		report(kv.Key, kv.Value)
	})
	fmt.Fprintf(w, "\n}\n")
}

func StartExpvarServer(config ExpvarConfig) {
	if config.ExpvarEnabled {
		go func() {
			mux := http.NewServeMux()
			mux.HandleFunc(config.ExpvarPath, metricsHandler)
			http.ListenAndServe(":"+config.ExpvarPort, mux)
		}()
	}
}
