package collect

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
)

var namespace = "node"

//Define a struct for you collector that contains pointers
//to prometheus descriptors for each metric you wish to expose.
//Note you can also include fields of other types if they provide utility
//but we just won't be exposing them as metrics.
type loadavgCollector struct {
	metrics []typedDesc
}

type typedDesc struct {
	desc      *prometheus.Desc
	valueType prometheus.ValueType
}

//You must create a constructor for you collector that
//initializes every descriptor and returns a pointer to the collector
func NewloadavgCollector() *loadavgCollector {
	return &loadavgCollector{
		metrics: []typedDesc{
			{prometheus.NewDesc(namespace+"_load1", "1m load average.", nil, nil), prometheus.GaugeValue},
			{prometheus.NewDesc(namespace+"_load5", "5m load average.", []string{"5min"}, nil), prometheus.GaugeValue},
			{prometheus.NewDesc(namespace+"_load15", "15m load average.", nil, nil), prometheus.GaugeValue},
		},
	}
}

//Each and every collector must implement the Describe function.
//It essentially writes all descriptors to the prometheus desc channel.
func (collector *loadavgCollector) Describe(ch chan<- *prometheus.Desc) {

	//Update this section with the each metric you create for a given collector
	ch <- collector.metrics[1].desc
}

//Collect implements required collect function for all promehteus collectors
func (collector *loadavgCollector) Collect(ch chan<- prometheus.Metric) {

	//Implement logic here to determine proper metric value to return to prometheus
	//for each descriptor or call other functions that do so.
	loads, err := GetLoad()
	if err != nil {
		log.Print("get loadavg error: ", err)
	}

	//Write latest value for each metric in the prometheus metric channel.
	//Note that you can pass CounterValue, GaugeValue, or UntypedValue types here.

	for i, load := range loads {
		ch <- prometheus.MustNewConstMetric(collector.metrics[i].desc, prometheus.GaugeValue, load)
	}

}
