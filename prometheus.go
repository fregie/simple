package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func runPromHttp(addr string) error {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	s := &http.Server{
		Addr:    addr,
		Handler: mux,
	}
	registerSessionGauge()
	return s.ListenAndServe()
}

func registerSessionGauge() {
	promauto.NewGaugeFunc(prometheus.GaugeOpts{
		Namespace: "simple",
		Subsystem: "session",
		Name:      "total",
	}, func() float64 {
		sessions := sessManager.GetAllSession()
		return float64(len(sessions))
	})
	protos := sessManager.GetProtos()
	for _, proto := range protos {
		p := proto
		promauto.NewGaugeFunc(prometheus.GaugeOpts{
			Namespace: "simple",
			Subsystem: "session",
			Name:      p,
		}, func() float64 {
			sessions, err := sessManager.GetProtoSessions(p)
			if err != nil {
				return 0
			}
			return float64(len(sessions))
		})
	}
}
