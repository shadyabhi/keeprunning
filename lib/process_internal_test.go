package keeprunning

import (
	"os"
	"strings"
	"testing"
	"time"
)

func TestIsRestart(t *testing.T) {

	var testCases = []struct {
		t         time.Time
		maxUptime time.Duration
		out       bool
	}{
		{time.Now().Add(time.Duration(0 * time.Second)), time.Duration(1 * time.Second), false},
		{time.Now().Add(time.Duration(1 * time.Second)), time.Duration(2 * time.Second), false},
	}

	for i, tt := range testCases {
		ret := isRestart(tt.t, tt.maxUptime)
		if ret != tt.out {
			t.Errorf("%d: Assertion failed. t: %s, maxUptime: %s, expected: %t. ", i, tt.t, tt.maxUptime, tt.out)
		}
	}
}

func TestStartProc(t *testing.T) {
	_, err := startProc(strings.Split("echo hello", " "), os.Stdout)
	if err != nil {
		t.Errorf("Error executing getting output of command: %s", err)
	}
}
