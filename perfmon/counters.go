package perfmon

import (
	"bytes"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"sync"

	. "github.com/intelsdi-x/snap-plugin-utilities/logger"
)

/*
 * Takes in a string of metric names requested in task
 * Returns a map of metric names to their values
 */
func GetPowershellData(mts []string) map[string]float64 {
	LogDebug("Begin gathering metric data from system", "metric_count", len(mts))

	runtime.GOMAXPROCS(runtime.NumCPU())
	// Map to store all the metrics with their values to pass to perfmon.go
	metricValues := make(map[string]float64)
	var mutex = &sync.Mutex{} // will synchronize access to to shared state (metricValues) across multiple goroutines

	// Have powershell command available for each metric name
	argValues := map[string]string{
		"processor_percent_time":    "(get-counter -Counter \"\\Processor(_Total)\\% Processor Time\" -ErrorAction Stop).CounterSamples.CookedValue",
		"logicalDisk_free_space":    "(get-counter -Counter \"\\LogicalDisk(_Total)\\% Free Space\" -ErrorAction Stop).CounterSamples.CookedValue",
		"physicalDisk_idle_time":    "(get-counter -Counter \"\\PhysicalDisk(_total)\\% Idle Time\" -ErrorAction Stop).CounterSamples.CookedValue",
		"physicalDisk_avg_read":     "(get-counter -Counter \"\\PhysicalDisk(_total)\\Avg. Disk sec/Read\" -ErrorAction Stop).CounterSamples.CookedValue",
		"physicalDisk_avg_write":    "(get-counter -Counter \"\\PhysicalDisk(_total)\\Avg. Disk sec/Write\" -ErrorAction Stop).CounterSamples.CookedValue",
		"physicalDisk_queue_length": "(get-counter -Counter \"\\PhysicalDisk(_total)\\Current Disk Queue Length\" -ErrorAction Stop).CounterSamples.CookedValue",
		"memory_committed_bytes":    "(get-counter -Counter \"\\Memory\\% Committed Bytes in Use\" -ErrorAction Stop).CounterSamples.CookedValue",
		"memory_available_mbytes":   "(get-counter -Counter \"\\Memory\\Available MBytes\" -ErrorAction Stop).CounterSamples.CookedValue",
		"memory_pagespersec":        "(get-counter -Counter \"\\Memory\\Pages/sec\" -ErrorAction Stop).CounterSamples.CookedValue",
		"pagingFile_percent_usage":  "(get-counter -Counter \"\\Paging File(_total)\\% Usage\" -ErrorAction Stop).CounterSamples.CookedValue",
		"system_up_time":            "(get-counter -Counter \"\\System\\System Up Time\" -ErrorAction Stop).CounterSamples.CookedValue",
		"system_context_switches":   "(get-counter -Counter \"\\System\\Context Switches/sec\" -ErrorAction Stop).CounterSamples.CookedValue"}
	cmdName := "powershell"
	var wg sync.WaitGroup

	// For each metric the user has requested, wait for responses (goroutines)
	wg.Add(len(mts))

	// Get data for each metric requested concurrently
	for _, metricName := range mts {
		go func(metricName string) {
			defer wg.Done() // defer pushes function call onto a list. Function is executed after surrounding function (goroutine) returns.
			// Command() returns a Cmd struct to execute named program with args, which is then executed by Run() further down
			cmdArg := argValues[metricName]
			cmd := exec.Command(cmdName, cmdArg)
			// Buffer is a variable-sized buffer of bytes with Read and Write methods; needs no initialization
			var counterOut bytes.Buffer
			var stderr bytes.Buffer
			// Stdout and Stderr of exec package specify processes' standard output and error channels
			cmd.Stdout = &counterOut
			cmd.Stderr = &stderr
			// Run() starts the command and waits for it to complete; typically returns error as type *ExitError - this doesn't provide sufficient error detail, so I use Stderr property of Command object as well
			err := cmd.Run()
			// If there is an error with command execution, log it, but keep going to get all the other metric values (don't return)
			if err != nil {
				LogError("Failed to execute command", "error", stderr.String())
			}

			// counterOut.String() adds a newline for some reason, so it must be removed first
			metricValue, formatErr := strconv.ParseFloat(strings.TrimSpace(counterOut.String()), 64)
			// Check to see if there was an error in parsing the value (this could happen if there are multiple values returned(doing (*) instead of (_total)), if no values are returned, or if the counter cannot be found on the system)
			if formatErr != nil {
				errorMessage := "There was an error with " + metricName
				LogError(errorMessage, "error", formatErr)
				metricValue = -1
			}
			mutex.Lock()
			metricValues[metricName] = metricValue
			mutex.Unlock()
		}(metricName)
	}

	wg.Wait()

	// Return map of requested metrics and their values
	return metricValues

}
