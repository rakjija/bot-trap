package metrics

import "github.com/prometheus/client_golang/prometheus"

var registered = false

func Init() {
	if registered {
		return
	}
	registered = true

	prometheus.MustRegister(LogSaveCounter)
	prometheus.MustRegister(LogErrorCounter)
	prometheus.MustRegister(SuspiciousLogCounter)
	prometheus.MustRegister(HttpDuration)
}
