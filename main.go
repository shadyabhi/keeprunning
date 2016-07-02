package main

import (
	"log"
	"os"

	keeprunning "github.com/shadyabhi/keeprunning/lib"
)

func main() {
	cArgs, err := initArgs()
	if err != nil {
		log.Fatalf("Error parsing command-line arguments: %s", err)
	}

	keeprunning.StartProcAndMonitor(cArgs.Cmd, cArgs.MaxUptime, os.Stdout)
}
