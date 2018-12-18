package monitor

import (
	"fmt"
	"github.com/DSiSc/craft/log"
	"net/http"
	_ "net/http/pprof"
)

type PprofConfig struct {
	PprofEnabled bool
	PprofPort    string
}

func StartPprofServer(config PprofConfig) {
	if config.PprofEnabled {
		go func() {
			log.Info(fmt.Sprintf("pprof: %x", http.ListenAndServe(":"+config.PprofPort, nil)))
		}()
	}
}
