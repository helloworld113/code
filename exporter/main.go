package main

import (
	"log"
	"myexporter/collect"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	prometheus.MustRegister(collect.NewloadavgCollector())

	http.Handle("/metrics", promhttp.Handler())
	log.Print("expose /metrics use port :8085")
	log.Fatal(http.ListenAndServe(":8085", nil))
}
