package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	LogSaveCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "log_save_total",
			Help: "총 저장된 로그 수",
		},
	)

	LogErrorCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "log_error_total",
			Help: "DB 저장 실패 횟수",
		},
	)

	SuspiciousLogCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "log_suspicious_total",
			Help: "의심되는 로그 메시지의 누적 횟수",
		},
	)
)
