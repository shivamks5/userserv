package metrics

import (
	"github.com/go-kit/kit/metrics"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	RequestCount   metrics.Counter   // total number of requests per method
	RequestLatency metrics.Histogram // how long each request takes (in microseconds)
	RequestErrors  metrics.Counter   // number of errors per method
}

func NewPrometheusMetrics(serviceName string) *Metrics {
	labels := []string{"method", "error"}
	return &Metrics{
		RequestCount: kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: serviceName,
			Subsystem: serviceName,
			Name:      "request_count",
			Help:      "Number of requests received",
		}, labels),
		RequestLatency: kitprometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: serviceName,
			Subsystem: serviceName,
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds",
		}, labels),
		RequestErrors: kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: serviceName,
			Subsystem: serviceName,
			Name:      "request_errors",
			Help:      "Number of errors encountered",
		}, labels),
	}
}
