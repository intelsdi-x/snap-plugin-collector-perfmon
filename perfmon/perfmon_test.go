package perfmon

import (
	"testing"
	"time"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPerfmonCollector(t *testing.T) {
    pm := PerfmonCollector{}

    Convey("Test PermonCollector", t, func() {
        Convey("Collect physicalDisk_idle_time", func() {
            metrics := []plugin.Metric{
                // Create fake physicalDisk_idle_time metric to make sure the CollectMetrics function is functioning correctly and returns how it should
                plugin.Metric {
                    Namespace:  plugin.NewNamespace("intel", "perfmon", "physicalDisk_idle_time"),
                    Config:     map[string]interface{}{"testfloat": float64(100.0211508)},
                    Data:       100.0211508,
                    Unit:       "float",        
                    Timestamp:  time.Now(),
                },
            }
            mts, err := pm.CollectMetrics(metrics)
            So(mts, ShouldNotBeEmpty)
            So(err, ShouldBeNil)
            So(mts[0].Data, ShouldEqual, 100.0211508)
        })
        Convey("Collect physicalDisk_avg_read", func() {
            metrics := []plugin.Metric{
                // Create fake physicalDisk_avg_read metric to make sure the CollectMetrics function is functioning correctly and returns how it should
                plugin.Metric {
                    Namespace:  plugin.NewNamespace("intel", "perfmon", "physicalDisk_avg_read"),
                    Config:     map[string]interface{}{"testfloat": float64(0)},
                    Data:       0,
                    Unit:       "float",        
                    Timestamp:  time.Now(),
                },
            }
            mts, err := pm.CollectMetrics(metrics)
            So(mts, ShouldNotBeEmpty)
            So(err, ShouldBeNil)
            So(mts[0].Data, ShouldEqual, 0)
        })
        Convey("Collect physicalDisk_avg_write", func() {
            metrics := []plugin.Metric{
                // Create fake physicalDisk_avg_write metric to make sure the CollectMetrics function is functioning correctly and returns how it should
                plugin.Metric {
                    Namespace:  plugin.NewNamespace("intel", "perfmon", "physicalDisk_avg_write"),
                    Config:     map[string]interface{}{"testfloat": float64(0.00077732235)},
                    Data:       0.00077732235,
                    Unit:       "float",        
                    Timestamp:  time.Now(),
                },
            }
            mts, err := pm.CollectMetrics(metrics)
            So(mts, ShouldNotBeEmpty)
            So(err, ShouldBeNil)
            So(mts[0].Data, ShouldEqual, 0.00077732235)
        })
        Convey("Collect physicalDisk_queue_length", func() {
            metrics := []plugin.Metric{
                // Create fake physicalDisk_queue_length metric to make sure the CollectMetrics function is functioning correctly and returns how it should
                plugin.Metric {
                    Namespace:  plugin.NewNamespace("intel", "perfmon", "physicalDisk_queue_length"),
                    Config:     map[string]interface{}{"testfloat": float64(1)},
                    Data:       1,
                    Unit:       "float",        
                    Timestamp:  time.Now(),
                },
            }
            mts, err := pm.CollectMetrics(metrics)
            So(mts, ShouldNotBeEmpty)
            So(err, ShouldBeNil)
            So(mts[0].Data, ShouldEqual, 1)
        })
        Convey("Collect memory_available", func() {
            metrics := []plugin.Metric{
                // Create fake memory_available metric to make sure the CollectMetrics function is functioning correctly and returns how it should
                plugin.Metric {
                    Namespace:  plugin.NewNamespace("intel", "perfmon", "memory_available_mbytes"),
                    Config:     map[string]interface{}{"testfloat": float64(10656)},
                    Data:       10656,
                    Unit:       "float",        
                    Timestamp:  time.Now(),
                },
            }
            mts, err := pm.CollectMetrics(metrics)
            So(mts, ShouldNotBeEmpty)
            So(err, ShouldBeNil)
            So(mts[0].Data, ShouldEqual, 10656)
        })
        Convey("Collect memory_committed", func() {
            metrics := []plugin.Metric{
                // Create fake memory_committed metric to make sure the CollectMetrics function is functioning correctly and returns how it should
                plugin.Metric {
                    Namespace:  plugin.NewNamespace("intel", "perfmon", "memory_committed_bytes"),
                    Config:     map[string]interface{}{"testfloat": float64(36.37)},
                    Data:       36.37,
                    Unit:       "float",        
                    Timestamp:  time.Now(),
                },
            }
            mts, err := pm.CollectMetrics(metrics)
            So(mts, ShouldNotBeEmpty)
            So(err, ShouldBeNil)
            So(mts[0].Data, ShouldEqual, 36.37)
        })
        Convey("Collect memory_pagespersec", func() {
            metrics := []plugin.Metric{
                // Create fake memory_pagespersec metric to make sure the CollectMetrics function is functioning correctly and returns how it should
                plugin.Metric {
                    Namespace:  plugin.NewNamespace("intel", "perfmon", "memory_pagespersec"),
                    Config:     map[string]interface{}{"testfloat": float64(0)},
                    Data:       0,
                    Unit:       "float",        
                    Timestamp:  time.Now(),
                },
            }
            mts, err := pm.CollectMetrics(metrics)
            So(mts, ShouldNotBeEmpty)
            So(err, ShouldBeNil)
            So(mts[0].Data, ShouldEqual, 0)
        })
        Convey("Collect page_usage", func() {
            metrics := []plugin.Metric{
                // Create fake page_usage metric to make sure the CollectMetrics function is functioning correctly and returns how it should
                plugin.Metric {
                    Namespace:  plugin.NewNamespace("intel", "perfmon", "page_usage"),
                    Config:     map[string]interface{}{"testfloat": float64(1.42)},
                    Data:       1.42,
                    Unit:       "float",        
                    Timestamp:  time.Now(),
                },
            }
            mts, err := pm.CollectMetrics(metrics)
            So(mts, ShouldNotBeEmpty)
            So(err, ShouldBeNil)
            So(mts[0].Data, ShouldEqual, 1.42)
        })
        Convey("Collect system_up_time", func() {
            metrics := []plugin.Metric{
                // Create fake system_up_time metric to make sure the CollectMetrics function is functioning correctly and returns how it should
                plugin.Metric {
                    Namespace:  plugin.NewNamespace("intel", "perfmon", "system_up_time"),
                    Config:     map[string]interface{}{"testfloat": float64(668074.8282)},
                    Data:       668074.8282,
                    Unit:       "float",        
                    Timestamp:  time.Now(),
                },
            }
            mts, err := pm.CollectMetrics(metrics)
            So(mts, ShouldNotBeEmpty)
            So(err, ShouldBeNil)
            So(mts[0].Data, ShouldEqual, 668074.8282)
        })
        Convey("Collect system_context_switches", func() {
            metrics := []plugin.Metric{
                // Create fake system_context_switches metric to make sure the CollectMetrics function is functioning correctly and returns how it should
                plugin.Metric {
                    Namespace:  plugin.NewNamespace("intel", "perfmon", "system_context_switches"),
                    Config:     map[string]interface{}{"testfloat": float64(24768.5784167836)},
                    Data:       24768.5784167836,
                    Unit:       "float",        
                    Timestamp:  time.Now(),
                },
            }
            mts, err := pm.CollectMetrics(metrics)
            So(mts, ShouldNotBeEmpty)
            So(err, ShouldBeNil)
            So(mts[0].Data, ShouldEqual, 24768.5784167836)
        })
        Convey("Collect processor_time", func() {
            metrics := []plugin.Metric{
                // Create fake processor_time metric to make sure the CollectMetrics function is functioning correctly and returns how it should
                plugin.Metric {
                    Namespace:  plugin.NewNamespace("intel", "perfmon", "processor_time"),
                    Config:     map[string]interface{}{"testfloat": float64(0.78)},
                    Data:       0.78,
                    Unit:       "float",        
                    Timestamp:  time.Now(),
                },
            }
            mts, err := pm.CollectMetrics(metrics)
            So(mts, ShouldNotBeEmpty)
            So(err, ShouldBeNil)
            So(mts[0].Data, ShouldEqual, 0.78)
        })
        Convey("Collect logical_disk_free", func() {
            metrics := []plugin.Metric{
                // Create fake logical_disk_free metric to make sure the CollectMetrics function is functioning correctly and returns how it should
                plugin.Metric {
                    Namespace:  plugin.NewNamespace("intel", "perfmon", "logical_disk_free"),
                    Config:     map[string]interface{}{"testfloat": float64(50.724595)},
                    Data:       50.724595,
                    Unit:       "float",        
                    Timestamp:  time.Now(),
                },
            }
            mts, err := pm.CollectMetrics(metrics)
            So(mts, ShouldNotBeEmpty)
            So(err, ShouldBeNil)
            So(mts[0].Data, ShouldEqual, 50.724595)
        })
    })

    Convey("Test GetMetricTypes", t, func() {
		pm := PerfmonCollector{}

		Convey("Collect All Metrics String", func() {
			mt, err := pm.GetMetricTypes(nil)
			So(err, ShouldBeNil)
			So(len(mt), ShouldEqual, 12)
		})
	})

    Convey("Test GetConfigPolicy", t, func() {
		pm := PerfmonCollector{}
		_, err := pm.GetConfigPolicy()

		Convey("No error returned", func() {
			So(err, ShouldBeNil)
		})
	})

}