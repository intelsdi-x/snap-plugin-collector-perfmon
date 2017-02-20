# Snap collector plugin - Perfmon
This plugin collects metrics from Windows counters (which the Perfmon GUI accesses) which allows us to gather system information such as system up time, CPU utilization, memory available, etc.

It's used in the [Snap framework](http://github.com:intelsdi-x/snap).

1. [Getting Started](#getting-started)
  * [System Requirements](#system-requirements)
  * [Installation](#installation)
  * [Configuration and Usage](#configuration-and-usage)
2. [Documentation](#documentation)
  * [Collected Metrics](#collected-metrics)
  * [Examples](#examples)
  * [Roadmap](#roadmap)
3. [Community Support](#community-support)
4. [Contributing](#contributing)
5. [Known Issues](#known-issues)
6. [License](#license-and-authors)
7. [Acknowledgements](#acknowledgements)

## Getting Started
### System Requirements 
* [golang 1.7+](https://golang.org/dl/) (needed only for building as code is written in Go)

### Operating systems
All OSs currently supported by this plugin:
* Currently tested on Windows 10, but should run on any Windows system 

### Installation
#### Download perfmon plugin binary:
You can get the pre-built binaries under the plugin's [release](https://github.com/Snap-for-Windows/snap-plugin-collector-perfmon/releases) page.  For Snap, check [here](https://github.com/intelsdi-x/snap/releases).


#### To build the plugin binary:
Need to create a build script for this plugin still.

### Configuration and Usage
* Set up the [Snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started)

## Documentation
There are a number of other resources you can review to learn to use this plugin:

* [Snap perfmon unit test](https://github.com/Snap-for-Windows/snap-plugin-collector-perfmon/perfmon/perfmon_test.go)
* [Snap perfmon examples](#examples)
* [Possible add-on counters and descriptions](https://blogs.technet.microsoft.com/askcore/2012/03/16/windows-performance-monitor-disk-counters-explained/)
* [Common counters and thresholds](https://blogs.technet.microsoft.com/bulentozkir/2014/02/14/top-10-most-important-performance-counters-for-windows-and-their-recommended-values/)
* [More Common counters and thresholds](https://support.symantec.com/en_US/article.HOWTO9722.html)
* [Even More common counters and thresholds](http://techgenix.com/Key-Performance-Monitor-Counters/)

### Collected Metrics
Currently, this plugin has the ability to gather the following metrics:

Namespace | Description (optional)
----------|-----------------------
/intel/perfmon/physicalDisk_idle_time | percent time that hard disk is idle over a measurement interval
/intel/perfmon/physicalDisk_avg_read | average time (in seconds) of a read of data from the disk
/intel/perfmon/physicalDisk_avg_write | average time (in seconds) of a write of data to the disk
/intel/perfmon/physicalDisk_queue_length | average number of both read and write requests that were queued for all disks during sample interval
/intel/perfmon/memory_committed_bytes | bytes of RAM being used
/intel/perfmon/memory_available_mbytes | mbytes of RAM available for use
/intel/perfmon/memory_pagespersec | rate at which pages are read from or written to disk 
/intel/perfmon/page_usage | percentage of paging file (for virtual memory) being used
/intel/perfmon/system_up_time | seconds since server last rebooted
/intel/perfmon/system_context_switches | how frequently the processor has to switch from user- to kernel-mode per second
/intel/perfmon/processor_time | percentage of elapsed time that all of process threads used the processor (in this case, all the processors) to execute instructions
/intel/perfmon/logical_disk_free | percentage of the total usable space on the selected logical disk that is free (in this case, the total of all logical disks on machine)

### Examples
This is an example running perfmon and writing data to a file. It is assumed that you are using the latest Snap binary and plugins.
It is also assumed that the user has a folder within the C: drive called "SnapLogs".

The example is run from a directory which includes snaptel, snapteld, along with the plugins and task file.

In one terminal window, open the Snap daemon (in this case with logging set to 1 and trust disabled):
```
$ snapteld -l 1 -t 0
```

In another terminal window:
Load perfmon plugin
```
$ snaptel plugin load snap-plugin-collector-perfmon
Plugin loaded
Name: perfmon-collector
Version: 1
Type: collector
Signed: false
Loaded Time: Mon, 20 Feb 2017 11:17:17 MST
```
See available metrics for your system
```
$ snaptel metric list
```

Create a task manifest file (e.g. `task-perfmon.json`):    
```json
{ 
    "version": 1,
    "schedule": {
        "type": "simple",
        "interval": "30s"
    },
    "max-failures": 10,
    "workflow": {
        "collect": {
            "metrics": {
                "/intel/perfmon/memory_committed_bytes": {},
                "/intel/perfmon/memory_available_mbytes": {},
                "/intel/perfmon/processor_time": {}
            },
            "process": [
                {
                    "plugin_name": "passthru-grpc",
                    "process": null,
                    "publish": [
                        {
                            "plugin_name": "mock-file-grpc",
                            "config": {
                                "file": "C:\\SnapLogs\\perfmon_published_revised.log"
                            }
                        }
                    ]
                }
            ]
        }
    }
}
```
Load passthru plugin for processing:
```
$ snaptel plugin load snap-plugin-processor-passthru-grpc
Plugin loaded
Name: passthru-grpc
Version: 1
Type: processor
Signed: false
Loaded Time: Mon, 20 Feb 2017 11:16:37 MST
```

Load file plugin for publishing:
```
$ snaptel plugin load snap-plugin-publisher-mock-file-grpc
Plugin loaded
Name: mock-file-grpc
Version: 1
Type: publisher
Signed: false
Loaded Time: Mon, 20 Feb 2017 11:16:58 MST
```

Create task:
```
$ snaptel task create -t task-perfmon.json
Using task manifest to create task
Task created
ID: 4a156b0f-582f-4a13-8d67-120a2ba72e1d
Name: Task-4a156b0f-582f-4a13-8d67-120a2ba72e1d
State: Running
```

See file output (this is just part of the file):
```
2017-02-20 12:05:57.6877987 -0700 MST|[{intel  } {perfmon  } {processor_time  }]|0.658707496758626|tags[plugin_running_on:DESKTOP-0GETRGO]
2017-02-20 12:05:57.6877987 -0700 MST|[{intel  } {perfmon  } {memory_available_mbytes  }]|10381|tags[plugin_running_on:DESKTOP-0GETRGO]
2017-02-20 12:05:57.6877987 -0700 MST|[{intel  } {perfmon  } {memory_committed_bytes  }]|38.7844268555717|tags[plugin_running_on:DESKTOP-0GETRGO]
2017-02-20 12:06:27.6945658 -0700 MST|[{intel  } {perfmon  } {processor_time  }]|1.25933460733828|tags[plugin_running_on:DESKTOP-0GETRGO]
2017-02-20 12:06:27.6945658 -0700 MST|[{intel  } {perfmon  } {memory_available_mbytes  }]|10382|tags[plugin_running_on:DESKTOP-0GETRGO]
2017-02-20 12:06:27.6945658 -0700 MST|[{intel  } {perfmon  } {memory_committed_bytes  }]|38.8447365115501|tags[plugin_running_on:DESKTOP-0GETRGO]
2017-02-20 12:06:57.6960669 -0700 MST|[{intel  } {perfmon  } {processor_time  }]|1.84033283218797|tags[plugin_running_on:DESKTOP-0GETRGO]
2017-02-20 12:06:57.6960669 -0700 MST|[{intel  } {perfmon  } {memory_available_mbytes  }]|10381|tags[plugin_running_on:DESKTOP-0GETRGO]
2017-02-20 12:06:57.6960669 -0700 MST|[{intel  } {perfmon  } {memory_committed_bytes  }]|38.8223214631021|tags[plugin_running_on:DESKTOP-0GETRGO]
```

Stop task:
```
$ snaptel task stop 4a156b0f-582f-4a13-8d67-120a2ba72e1d
Task stopped:
ID: 4a156b0f-582f-4a13-8d67-120a2ba72e1d
```

### Roadmap
There isn't a current roadmap for this plugin, but it is in active development. As we launch this plugin, we do not have any outstanding requirements for the next release. If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-collector-perfmon/issues/new) and/or submit a [pull request](https://github.com/intelsdi-x/snap-plugin-collector-perfmon/pulls).

## Community Support
This repository is one of **many** plugins in **Snap**, a powerful telemetry framework. See the full project at http://github.com/intelsdi-x/snap To reach out to other users, head to the [main framework](https://github.com/intelsdi-x/snap#community-support)

## Contributing
We love contributions!

There's more than one way to give back, from examples to blogs to code updates. See our recommended process in [CONTRIBUTING.md](CONTRIBUTING.md).

## License
[Snap](http://github.com:intelsdi-x/snap), along with this plugin, is an Open Source software released under the Apache 2.0 [License](LICENSE).

## Acknowledgements
* Author: [@mathewlk](https://github.com/mathewlk/)

And **thank you!** Your contribution, through code and participation, is incredibly important to us.