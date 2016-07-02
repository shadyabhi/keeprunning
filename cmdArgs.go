package main

import (
	"flag"
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type cmdArgs struct {
	Cmd       string
	MaxUptime time.Duration
}

func initArgs() (cmd cmdArgs, err error) {
	cArgs := cmdArgs{}

	flag.StringVar(&cArgs.Cmd, "cmd", "", "Path to external binary that will be executed/monitored")
	uptime := flag.Int("uptime", 0, "Restart process after x seconds. (Default: 0, don't restart). x=<0 means don't try restarting. ")
	flag.Parse()

	cArgs.MaxUptime = time.Duration(*uptime) * time.Second
	if err := checkArgs(cArgs); err != nil {
		return cmdArgs{}, err
	}
	return cArgs, nil
}

func checkArgs(cArgs cmdArgs) error {
	cmd := strings.Split(cArgs.Cmd, " ")[0]
	file, err := os.Stat(cmd)
	if err != nil {
		return errors.Wrap(err, "File location provided is not valid")
	}
	if !file.Mode().IsRegular() {
		return errors.Wrap(errors.New("Supplied file is not a regular file"), "")
	}
	return nil
}
