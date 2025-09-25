package metric

import "github.com/prometheus/client_golang/prometheus"

const (
	namespaceError = "error"
	namespaceSync  = "synchronizer"
)

var (
	// Errors errors count metric.
	Errors = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespaceError,
			Name:      "errors",
			Help:      "",
		}, []string{"error"})

	// Reorgs block reorg count
	Reorgs = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: namespaceSync,
			Name:      "reorgs",
			Help:      "",
		})

	// LastBlockNum last block synced
	LastBlockNum = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: namespaceSync,
			Name:      "synced_last_block_num",
			Help:      "",
		})

	// EthLastBlockNum last eth block synced
	EthLastBlockNum = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: namespaceSync,
			Name:      "eth_last_block_num",
			Help:      "",
		})
	// LastBatchNum last batch synced
	LastBatchNum = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: namespaceSync,
			Name:      "synced_last_batch_num",
			Help:      "",
		})

	// EthLastBatchNum last eth batch synced
	EthLastBatchNum = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: namespaceSync,
			Name:      "eth_last_batch_num",
			Help:      "",
		})
)

// CollectError collect the error message and increment
// the error count
func CollectError(err error) {
	Errors.With(map[string]string{"error": err.Error()}).Inc()
}
