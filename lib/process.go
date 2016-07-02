package keeprunning

import (
	"io"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/pkg/errors"
)

func startProc(cmdPath []string, output io.Writer) (*exec.Cmd, error) {
	c := exec.Command(cmdPath[0], cmdPath[1:]...)
	c.Stdout = output

	if err := c.Start(); err != nil {
		// Backoff and try again in some time
		sleepT := 1 * time.Second
		log.Printf("Sleeping for %d and then retry. ")
		time.Sleep(sleepT)
		return nil, errors.Wrap(err, "Error starting process. Will sleep and retry")
	}
	return c, nil

}

func isRestart(t time.Time, maxUptime time.Duration) bool {
	if maxUptime == 0 {
		return false
	}

	if time.Since(t) >= maxUptime {
		return true
	}
	return false
}

func processWatchdog(cmd *exec.Cmd, startTime time.Time, maxUptime time.Duration) {
	if maxUptime == time.Duration(0) {
		log.Printf("Started process with pid: %d, running indefinitely", cmd.Process.Pid)
	} else {
		log.Printf("Started process with pid: %d, killing after %s", cmd.Process.Pid, maxUptime)
	}

	for {
		if isRestart(startTime, maxUptime) {
			err := cmd.Process.Kill()
			if err != nil {
				log.Printf("Error killing process: %s", err)
			} else {
				// Get out of here
				break
			}
		}

		time.Sleep(1 * time.Second)
	}
}

// StartProcAndMonitor starts process and monitors it.
// cmdPath: Path to the binary
// maxUptime: Restart after x seconds. if x=0, never restart, x<0, don't try restart
func StartProcAndMonitor(cmdPath string, maxUptime time.Duration, output io.Writer) {
	cmd := strings.Split(cmdPath, " ")

	for {
		log.Printf("Starting process... : %s", cmdPath)
		execCmd, err := startProc(cmd, output)
		if err != nil {
			log.Printf("Error starting process. Backing off and trying again: %s", err)
			time.Sleep(1 * time.Second)
		}

		go processWatchdog(execCmd, time.Now(), maxUptime)

		if err := execCmd.Wait(); err != nil {
			log.Printf("Process died: %s", err)
		} else {
			log.Printf("Process exited gracefully.")
		}

		if maxUptime < 0 {
			break
		}
	}
}
