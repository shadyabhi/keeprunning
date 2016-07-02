[![Build Status](https://travis-ci.org/shadyabhi/keeprunning.png)](https://travis-ci.org/shadyabhi/keeprunning)
[![codecov](https://codecov.io/gh/shadyabhi/keeprunning/branch/master/graph/badge.svg)](https://codecov.io/gh/shadyabhi/keeprunning)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/shadyabhi/keeprunning)

# keeprunning

This makes sure that a process is always running on a system.

## Features

* Watches a process and makes sure that it restarts a process if it does.
* Periodic restarts of processses that are misbehaving. (optional)

## Installation

```
➜ $?=0 /Users/arastogi [ 4:29PM] % go get github.com/shadyabhi/keeprunning

>>>  6s elasped...
➜ $?=0 /Users/arastogi [ 4:29PM] % keeprunning -h
Usage of keeprunning:
  -cmd string
        Path to external binary that will be executed/monitored
  -uptime int
        Restart process after x seconds. (Default: 0, don't restart). x=<0 means don't try restarting.

>>>  0s elasped...
➜ $?=2 /Users/arastogi [ 4:29PM] %
```

## Examples

* This command runs "/bin/sleep 15" and starts monitoring that command. It will make sure that this command is run-over even if it gets killed by external factors.
```
keeprunning -cmd "/bin/sleep 15"
```

* Does the same as above but also restarts the command every 3600 seconds
```
keeprunning -cmd "/bin/sleep 3610" -uptime 3600
```

