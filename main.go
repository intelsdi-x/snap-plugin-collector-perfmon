package main

import (

	// Import the more recent gRPC (Go RPC plugin lib is now deprecated) snap plugin library
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin" 
	// Import our collector plugin implementation
	"github.com/Snap-for-Windows/snap-plugin-collector-perfmon/perfmon"
)

const (
	pluginName    = "perfmon-collector"
	pluginVersion = 1
)

//plugin bootstrap
func main() {
	plugin.StartCollector(
		perfmon.PerfmonCollector{},  
		pluginName,               
		pluginVersion)
}
