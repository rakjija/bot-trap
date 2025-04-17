package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var HttpDuration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "요청별 처리 시간 (초 단위)",
		Buckets: prometheus.DefBuckets,
	},
	[]string{"path", "method"},
)
