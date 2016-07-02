package keeprunning

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"testing"
)

func TestStartProcAndMonitor(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	buf := &bytes.Buffer{}
	StartProcAndMonitor("/bin/echo hello", -1, buf)
	s := bufio.NewScanner(buf)
	if s.Scan() {
		if s.Text() != "hello" {
			t.Errorf("Output was unexpected. Output: %s", s.Text())
		}
	}
}
