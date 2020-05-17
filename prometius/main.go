// package main

// import (
// 	"github.com/prometheus/client_golang/prometheus"
// )

// var (
// 	requestsIncoming = prometheus.NewCounterVec(
// 		prometheus.CounterOpts{
// 			Namespace: "etcd",
// 			Subsystem: "proxy",
// 			Name:      "requests_total",
// 			Help:      "Counter requests incoming by method.",
// 		}, []string{"method"})

// 	requestsHandled = prometheus.NewCounterVec(
// 		prometheus.CounterOpts{
// 			Namespace: "etcd",
// 			Subsystem: "proxy",
// 			Name:      "handled_total",
// 			Help:      "Counter of requests fully handled (by authoratitave servers)",
// 		}, []string{"method", "code"})
// )

// func main() {
// 	requestsIncoming.WithLabelValues("gelen").Inc()

// }
package main

import (
	"net/http"

	"log"
	"math/rand"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	counter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "golang",
			Name:      "my_counter",
			Help:      "This is my counter",
		})

	gauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "golang",
			Name:      "my_gauge",
			Help:      "This is my gauge",
		})

	histogram = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: "golang",
			Name:      "my_histogram",
			Help:      "This is my histogram",
		})

	summary = prometheus.NewSummary(
		prometheus.SummaryOpts{
			Namespace: "golang",
			Name:      "my_summary",
			Help:      "This is my summary",
		})
)

func main() {
	rand.Seed(time.Now().Unix())

	http.Handle("/metrics", promhttp.Handler())

	prometheus.MustRegister(counter)
	prometheus.MustRegister(gauge)
	prometheus.MustRegister(histogram)
	prometheus.MustRegister(summary)

	go func() {
		for {
			counter.Add(rand.Float64() * 5)
			gauge.Add(rand.Float64()*15 - 5)
			histogram.Observe(rand.Float64() * 10)
			summary.Observe(rand.Float64() * 10)

			time.Sleep(time.Second)
		}
	}()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
