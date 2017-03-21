package perfmon

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetCounters(t *testing.T) {
	Convey("Test GetPowershellData", t, func() {

		Convey("When metrics are valid", func() {
			counterDataMap := GetPowershellData(availableMetrics)

			Convey("Twelve non-negative counter values should be returned", func() {
				for _, val := range counterDataMap {
					So(val, ShouldBeGreaterThanOrEqualTo, 0)
				}
			})
		})

		// This test would be the same if the actual counter path given in counters.go was not found on the system
		Convey("When metric names are not valid", func() {
			metricNames := []string{"emory_available_mbytes", "memory_committed_bytes", "processor_percent_time"}
			counterDataMap := GetPowershellData(metricNames)

			Convey("One -1 counter value should be returned", func() {
				for key, val := range counterDataMap {
					if key == "emory_available" {
						So(val, ShouldEqual, -1)
					}
				}
			})
		})

	})

}
