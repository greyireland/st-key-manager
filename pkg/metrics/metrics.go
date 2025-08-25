package metrics

import (
	"fmt"
	"github.com/greyireland/log"
	"net/http"

	"github.com/VictoriaMetrics/metrics"
)

// Setup starts a dedicated metrics server at the given address.
// This function enables metrics reporting separate from pprof.
func Setup(address string) {
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		metrics.WritePrometheus(w, true)
	})
	//m.Handle("/debug/metrics", ExpHandler(metrics.DefaultRegistry))
	//m.Handle("/debug/metrics/prometheus2", promhttp.HandlerFor(prometheus2.DefaultGatherer, promhttp.HandlerOpts{
	//	EnableOpenMetrics: true,
	//}))
	log.Info("Starting metrics server", "addr", fmt.Sprintf("http://%s/metrics", address))
	go func() {
		if err := http.ListenAndServe(address, nil); err != nil {
			log.Error("Failure in running metrics server", "err", err)
		}
	}()
}
