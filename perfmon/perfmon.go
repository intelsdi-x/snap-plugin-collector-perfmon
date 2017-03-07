package perfmon

import (
	"fmt"
	"time"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
)

// Collector implementation. Needed even if empty. Create an empty struct to use as receiver type of methods below
type PerfmonCollector struct {
}

var availableMetrics = []string{
	"physicalDisk_idle_time", "physicalDisk_avg_read", "physicalDisk_avg_write", "physicalDisk_queue_length",
	"memory_committed_bytes", "memory_available_mbytes", "memory_pagespersec",
	"pagingFile_percent_usage",
	"system_up_time", "system_context_switches",
	"processor_percent_time",
	"logicalDisk_free_space"}

func stringInNamespace(givenString string) bool {
	for _, metricName := range availableMetrics {
		if metricName == givenString {
			return true
		}
	}
	return false
}

/*
* CollectMetrics collects metrics for testing.
* CollectMetrics() is called by Snap when a task (which is collecting one+ of the metrics returned from the GetMetricTypes()) is started.
* Input: A slice of all the metric types being collected.
* Output: A slice (list) of the collected metrics as plugin.Metric with their values and an error if failure.
 */
func (PerfmonCollector) CollectMetrics(mts []plugin.Metric) ([]plugin.Metric, error) {
	metrics := []plugin.Metric{} // Create a slice of MetricType objects. This is where the metrics requested by the task will be stored
	// Iterate through metrics first time to create slice of metric names to pass to GetPowershellData
	metricNames := make([]string, 0)
	for _, mt := range mts {
		fullNameSpace := mt.Namespace[2].Value + "_" + mt.Namespace[3].Value
		metricNames = append(metricNames, fullNameSpace)
	}

	// Get metric data from powershell script if data has not been set already (for testing). 0 means there was an error getting that metric from system
	counterData := GetPowershellData(metricNames)
	// Iterate through each of the metrics specified by task to collect
	for idx, mt := range mts {
		fullNameSpace := mt.Namespace[2].Value + "_" + mt.Namespace[3].Value
		// Make sure the metric given in the Task is actually a metric provided by this plugin
		if stringInNamespace(fullNameSpace) {
			mts[idx].Timestamp = time.Now() // Set metric timestamp
			// Make sure config hasn't been set for testing (SEE perfmon_test.go)
			if val, err := mt.Config.GetFloat("testfloat"); err == nil {
				mts[idx].Data = val
			} else {
				mts[idx].Data = counterData[fullNameSpace] // Set metric value
			}
			metrics = append(metrics, mts[idx])
		} else {
			return nil, fmt.Errorf("Invalid metric: %v", mt.Namespace.Strings())
		}
	}
	return metrics, nil
}

/*
 * GetMetricTypes returns a list of available metric types
 * GetMetricTypes() is called when this plugin is loaded in order to populate the "metric catalog" (where Snap
 * stores all of the available metrics for each plugin)
 * Input: Config info. This information comes from global Snap config settings
 * Output: A slice (list) of all plugin metrics, which are available to be collected by tasks
 */
func (PerfmonCollector) GetMetricTypes(cfg plugin.Config) ([]plugin.Metric, error) {
	// slice to store list of all available perfmon metrics
	mts := []plugin.Metric{}

	mts = append(mts, plugin.Metric{
		Namespace: plugin.NewNamespace("intel", "perfmon", "physicalDisk", "idle_time"),
		Version:   1,
	})
	mts = append(mts, plugin.Metric{
		Namespace: plugin.NewNamespace("intel", "perfmon", "physicalDisk", "avg_read"),
		Version:   1,
	})
	mts = append(mts, plugin.Metric{
		Namespace: plugin.NewNamespace("intel", "perfmon", "physicalDisk", "avg_write"),
		Version:   1,
	})
	mts = append(mts, plugin.Metric{
		Namespace: plugin.NewNamespace("intel", "perfmon", "physicalDisk", "queue_length"),
		Version:   1,
	})
	mts = append(mts, plugin.Metric{
		Namespace: plugin.NewNamespace("intel", "perfmon", "memory", "committed_bytes"),
		Version:   1,
	})
	mts = append(mts, plugin.Metric{
		Namespace: plugin.NewNamespace("intel", "perfmon", "memory", "available_mbytes"),
		Version:   1,
	})
	mts = append(mts, plugin.Metric{
		Namespace: plugin.NewNamespace("intel", "perfmon", "memory", "pagespersec"),
		Version:   1,
	})
	mts = append(mts, plugin.Metric{
		Namespace: plugin.NewNamespace("intel", "perfmon", "pagingFile", "percent_usage"),
		Version:   1,
	})
	mts = append(mts, plugin.Metric{
		Namespace: plugin.NewNamespace("intel", "perfmon", "system", "up_time"),
		Version:   1,
	})
	mts = append(mts, plugin.Metric{
		Namespace: plugin.NewNamespace("intel", "perfmon", "system", "context_switches"),
		Version:   1,
	})
	mts = append(mts, plugin.Metric{
		Namespace: plugin.NewNamespace("intel", "perfmon", "processor", "percent_time"),
		Version:   1,
	})
	mts = append(mts, plugin.Metric{
		Namespace: plugin.NewNamespace("intel", "perfmon", "logicalDisk", "free_space"),
		Version:   1,
	})

	return mts, nil
}

/*
 * GetConfigPolicy() returns the config policy for this plugin
 *   A config policy allows users to provide configuration info to the plugin and is provided in the task. Here we define what kind of config info this plugin can take and/or needs.
 */
func (PerfmonCollector) GetConfigPolicy() (plugin.ConfigPolicy, error) {
	policy := plugin.NewConfigPolicy()

	// This rule is simply for unit testing, so I can pass in my own values for each metric rather than getting them from counters.go, as the values are constantly changing in real time
	policy.AddNewFloatRule([]string{"random", "float"},
		"testfloat",
		false,
		plugin.SetMaxFloat(1000.0),
		plugin.SetMinFloat(0.0))

	// For now, assuming that perfmon has no configs. May need to add some if permissions becomes an issue.
	return *policy, nil
}
