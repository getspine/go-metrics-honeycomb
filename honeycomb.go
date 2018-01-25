package honeycomb

import (
	"time"

	"github.com/honeycombio/libhoney-go"
	"github.com/rcrowley/go-metrics"
)

type Reporter struct {
	Registry metrics.Registry
	Interval time.Duration
	WriteKey string
	Dataset  string

	stopped chan struct{}
}

func NewReporter(
	registry metrics.Registry,
	interval time.Duration,
	writeKey string,
	dataset string,
) *Reporter {
	r := &Reporter{
		Registry: registry,
		Interval: interval,
		WriteKey: writeKey,
		Dataset:  dataset,
	}
	r.Init()
	return r
}

func Honeycomb(
	registry metrics.Registry,
	interval time.Duration,
	writeKey string,
	dataset string,
) {
	NewReporter(registry, interval, writeKey, dataset).Run()
}

// Initializes the Honeycomb client.
func (r *Reporter) Init() {
	libhoney.Init(libhoney.Config{
		WriteKey: r.WriteKey,
		Dataset:  r.Dataset,
	})
}

// Convenience method around libhoney.AddField()
func (r *Reporter) AddField(key string, val interface{}) {
	libhoney.AddField(key, val)
}

// Blocks and starts reporting metrics from the provided registry to Honeycomb.
func (r *Reporter) Run() {
	defer r.Stop()
	for {
		select {
		case <-time.After(r.Interval):
			libhoney.SendNow(r.buildRequest())
		case <-r.stopped:
			return
		}
	}
}

// Stops the metrics reporting process and closes any connections to Honeycomb.
func (r *Reporter) Stop() {
	close(r.stopped)
	libhoney.Close()
}

func (r *Reporter) buildRequest() map[string]interface{} {
	metricsMap := make(map[string]interface{})
	r.Registry.Each(func(name string, metric interface{}) {
		metricsMap[name] = metric
	})
	return metricsMap
}
