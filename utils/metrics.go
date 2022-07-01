package utils

import (
	"base-app/config"

	"github.com/penglongli/gin-metrics/ginmetrics"
)

// exporter for prometheus metrics
func GetMetricsExporter() *ginmetrics.Monitor {

	m := ginmetrics.GetMonitor()

	// +optional set metric path, default /debug/metrics
	m.SetMetricPath(config.AppConfig.Server.MetricsPath)

	m.SetMetricPrefix(config.AppConfig.Server.MetricsPrefix)
	// +optional set slow time, default 5s
	m.SetSlowTime(10)
	// +optional set request duration, default {0.1, 0.3, 1.2, 5, 10}
	// used to p95, p99
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})

	return m

}
