// Documentation for keeprunning

/*
Package keeprunning providers helper function to start long running processes.

This makes sure that a process is always running on a system.

Features

- Watches a process and makes sure that it restarts a process if it does.

- Periodic restarts of processses that are misbehaving. (optional)


Usage

	➜ $?=0 /Users/arastogi/ % ./keeprunning -h
	Usage of ./keeprunning:
	  -cmd string
			Path to external binary that will be executed/monitored
	  -uptime int
			Restart process after x seconds. (Default: 0, don't restart). x=<0 means don't try restarting.

	>>>  1s elasped...
	➜ $?=0 /Users/arastogi [ 4:50PM] % keeprunning -cmd '/bin/sleep 15' -uptime 2
	2016/07/02 16:50:40 Starting process... : /bin/sleep 15
	2016/07/02 16:50:40 Started process with pid: 65561, killing after 2s
	2016/07/02 16:50:42 Process died: signal: killed
	2016/07/02 16:50:42 Starting process... : /bin/sleep 15
	2016/07/02 16:50:42 Started process with pid: 65562, killing after 2s
	2016/07/02 16:50:44 Process died: signal: killed
	2016/07/02 16:50:44 Starting process... : /bin/sleep 15
	2016/07/02 16:50:44 Started process with pid: 65563, killing after 2s
	^C

	>>>  4s elasped...
	➜ $?=130 /Users/arastogi [ 4:50PM] %

Examples

- This command runs "/bin/sleep 15" and starts monitoring that command. It will make sure that this command is run-over even if it gets killed by external factors.

	keeprunning -cmd "/bin/sleep 15"

- Does the same as above but also restarts the command every 3600 seconds

	keeprunning -cmd "/bin/sleep 3610" -uptime 3600
*/
package main
