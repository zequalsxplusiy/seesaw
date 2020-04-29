package metrics

import (
	"sync"
	"time"
)

var registerMetrics sync.Once

// LatencyMetric observes time taken for each ncc client call partitioned by action.
type LatencyMetric interface {
	Observe(action string, latency time.Duration)
}

// ResultMetric counts response codes partitioned by action.
type RequestMetric interface {
	Increment(action string)
}

var (
	// TimeSpent is the metric that ncc clients will update.
	RequestLatency LatencyMetric = noopLatency{}

	// RequestMetric is the metric to report how many types of a particular request is been called by
	// ncc client.
	Requests RequestMetric = noopResult{}
)

// Register registers metrics for NCC client. This can
// only be called once.
func Register(lm LatencyMetric, rm RequestMetric) {
	registerMetrics.Do(func() {
		RequestLatency = lm
		Requests = rm
	})
}


type noopLatency struct{}

func (noopLatency) Observe(string, time.Duration) {}

type noopResult struct{}

func (noopResult) Increment(string) {}
