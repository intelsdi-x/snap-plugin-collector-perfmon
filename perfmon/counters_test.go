package perfmon

import (
    "testing"
    "fmt"

    . "github.com/smartystreets/goconvey/convey"
)

func TestGetCounters(t *testing.T) {
    Convey("Test GetPowershellData", t, func() {

        Convey("When metrics are valid", func() {
            metricNames := []string{"physicalDisk_idle_time", "physicalDisk_avg_read", "physicalDisk_avg_write", "physicalDisk_queue_length",
            "memory_committed_bytes", "memory_available_mbytes", "memory_pagespersec",
            "page_usage", 
            "system_up_time", "system_context_switches", 
            "processor_time",
            "logical_disk_free"}
            counterDataMap := GetPowershellData(metricNames)

            Convey("Twelve non-negative counter values should be returned", func() {
                fmt.Println(counterDataMap)
                for _, val := range counterDataMap {
                    So(val, ShouldBeGreaterThanOrEqualTo, 0)
                }
            })
        })

        // This test would be the same if the actual counter path given in counters.go was not found on the system
        Convey("When metric names are not valid", func() {
            metricNames := []string{"emory_available_mbytes", "memory_committed_bytes", "processor_time"}
            counterDataMap := GetPowershellData(metricNames)

            Convey("One -1 counter value should be returned", func() {
                fmt.Println(counterDataMap)
                for key, val := range counterDataMap {
                    if key == "emory_available" {
                        So(val, ShouldEqual, -1)
                    }                
                }
            })
        })

    })

}