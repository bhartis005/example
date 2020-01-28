package main

import (
	"flag"

	log "github.com/sirupsen/logrus"
)

func init() {
	prometheus.RegisterMetric()
	// Register the summary and the histogram with Prometheus's default registry.
}

func main() {
	flag.Parse()

	// Periodically record some sample latencies for the three services.
	go func() {
		num := []int{2, 4, 6, 8, 10, 12, 14, 16, 13, 17, 90, 87}
		for _, number := range num {
			if number%2 == 0 {
				log.Info(number, "is even")
				prometheus.ErrorCounter.WithLabelValues("metric", "one").Inc()

			} else {
				log.Info(number, "is odd")
				prometheus.ErrorCounter.WithLabelValues("metric", "two").Inc()
			}
		}
	}()
	prometheus.ViewMetric()
}
