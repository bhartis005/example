package prom

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	//ErrorCounter is the counter initialization
	ErrorCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "error_counter",
			Help: "Total error count.",
		},
		[]string{"service", "error"},
	)
)

//RegisterMetric will register the metric
func RegisterMetric() {
	prometheus.MustRegister(ErrorCounter)
}

//ViewMetric will show the metrics
func ViewMetric() {
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
