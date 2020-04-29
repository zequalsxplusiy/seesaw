package metrics

import (
	"sync"
	"time"
)

var registerMetrics sync.Once

// LatencyMetric observes time taken for each ncc client call partitioned by action/task.
type LatencyMetric interface {
	Observe(task string, latency time.Duration)
}

var (
	// TimeSpent is the metric that ncc clients will update.
	RequestLatency LatencyMetric = noopLatency{}
)

// Register registers metrics for NCC client. This can
// only be called once.
func Register(m LatencyMetric) {
	registerMetrics.Do(func() {
		RequestLatency = m
	})
}


type noopLatency struct{}

func (noopLatency) Observe(string, time.Duration) {}
