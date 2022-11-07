DISCONTINUATION OF PROJECT. 

This project will no longer be maintained by Intel.

This project has been identified as having known security escapes.

Intel has ceased development and contributions including, but not limited to, maintenance, bug fixes, new releases, or updates, to this project.  

Intel no longer accepts patches to this project.

# DISCONTINUATION OF PROJECT 

**This project will no longer be maintained by Intel.  Intel will not provide or guarantee development of or support for this project, including but not limited to, maintenance, bug fixes, new releases or updates.  Patches to this project are no longer accepted by Intel. If you have an ongoing need to use this project, are interested in independently developing it, or would like to maintain patches for the community, please create your own fork of the project.**


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
* [powershell 4.0+ recommended](https://www.microsoft.com/en-us/download/details.aspx?id=40855): Powershell v2.0 does not have the Get-Counter powershell command, which this plugin requires. Windows Management Framework 4.0 or higher is recommended (although WMF 3.0 will work as well) as it includes an updated version of powershell. If the plugin fails, this is likely the cause.
* [golang 1.7+](https://golang.org/dl/): Needed only for building as code is written in Go.
* [glide 0.12.3+](http://glide.sh/): Required for developers in order to install correct package dependency versions.

### Operating systems
All OSs currently supported by this plugin:
* Currently tested on Windows 10, but should run on any Windows system 

### Installation
#### Download perfmon plugin binary:
You can get the pre-built binaries under the plugin's [release](https://github.com/intelsdi-x/snap-plugin-collector-perfmon/releases) page.  For Snap, check [here](https://github.com/intelsdi-x/snap/releases).


#### To build the plugin binary:
Build script for this plugin pending.  
For now, build manually:  
1. Download the plugin with `go get github.com/intelsdi-x/snap-plugin-collector-perfmon`
2. Navigate to the snap-plugin-collector-perfmon folder in your Go-Workspace
3. Use Glide to install correct dependency versions with `glide install`
4. Build the snap-plugin-collector-perfmon executable with `go install`
5. The plugin executable should now be located at [Path to Go-Workspace]\bin

### Configuration and Usage
* Set up the [Snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started)
* Due to current overhead issues with powershell, if gathering all plugin metrics, it is recommended that you use a task interval of 30 seconds or higher to ensure minimum failures when integrating with Snap.

## Documentation
There are a number of other resources you can review to learn to use this plugin:

* [Snap perfmon unit test](https://github.com/intelsdi-x/snap-plugin-collector-perfmon/perfmon/perfmon_test.go)
* [Snap perfmon examples](#examples)
* [Possible add-on counters and descriptions](https://blogs.technet.microsoft.com/askcore/2012/03/16/windows-performance-monitor-disk-counters-explained/)
* [Common counters and thresholds](https://blogs.technet.microsoft.com/bulentozkir/2014/02/14/top-10-most-important-performance-counters-for-windows-and-their-recommended-values/)
* [More Common counters and thresholds](https://support.symantec.com/en_US/article.HOWTO9722.html)
* [Even More common counters and thresholds](http://techgenix.com/Key-Performance-Monitor-Counters/)

### Collected Metrics
Currently, this plugin has the ability to gather the following metrics:

Namespace | Description (optional)
----------|-----------------------
/intel/perfmon/physicalDisk/idle_time | percent time that hard disk is idle over a measurement interval
/intel/perfmon/physicalDisk/avg_read | average time (in seconds) of a read of data from the disk
/intel/perfmon/physicalDisk/avg_write | average time (in seconds) of a write of data to the disk
/intel/perfmon/physicalDisk/queue_length | average number of both read and write requests that were queued for all disks during sample interval
/intel/perfmon/memory/committed_bytes | bytes of RAM being used
/intel/perfmon/memory/available_mbytes | mbytes of RAM available for use
/intel/perfmon/memory/pagespersec | rate at which pages are read from or written to disk 
/intel/perfmon/pagingFile/percent_usage | percentage of paging file (for virtual memory) being used
/intel/perfmon/system/up_time | seconds since server last rebooted
/intel/perfmon/system/context_switches | how frequently the processor has to switch from user- to kernel-mode per second
/intel/perfmon/processor/percent_time | percentage of elapsed time that all of process threads used the processor (in this case, all the processors) to execute instructions
/intel/perfmon/logicalDisk/free_space | percentage of the total usable space on the selected logical disk that is free (in this case, the total of all logical disks on machine)

### Examples
This is an example running perfmon and writing data to a file. It is assumed that you are using the latest Snap binary and plugins.
It is also assumed that the user has a folder within the C: drive called "SnapLogs".

The example is run from a directory which includes `snaptel`, `snapteld`, along with the plugins and task file.

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
                "/intel/perfmon/memory/committed_bytes": {},
                "/intel/perfmon/memory/available_mbytes": {},
                "/intel/perfmon/processor/percent_time": {}
            },
            "process": [
                {
                    "plugin_name": "passthru-grpc",
                    "process": null,
                    "publish": [
                        {
                            "plugin_name": "mock-file-grpc",
                            "config": {
                                "file": "C:\\SnapLogs\\perfmon.log"
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
