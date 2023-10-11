package main

import (
	"time"

	"github.com/dsandovale/refactored-journey/logging"
	"github.com/dsandovale/refactored-journey/random"
)

func main() {
	logger := logging.New(time.RFC3339, true)
	logger.Log("info", "starting up service")
	random.Calcule()
	logger.Log("warning", "no tasks found")
	logger.Log("error", "exiting: no work performed")
}
