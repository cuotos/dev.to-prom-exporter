package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"runtime"
)

var (
	Version   string
	GoVersion = runtime.Version()
)

func NewCollector(program string) prometheus.Collector {
	return prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Namespace: program,
			Name:      "build_info",
			Help:      fmt.Sprintf("A metric with constance '1' value labeled by versioni stil and goversion from which %s was built.", program),
			ConstLabels: prometheus.Labels{
				"version":   Version,
				"goversion": GoVersion,
			},
		},
		func() float64 { return 1 })
}
