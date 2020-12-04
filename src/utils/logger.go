package utils

import (
	"flag"
	"os"

	"github.com/google/logger"
)

// Init the logger on the app startup
func init() {
	flag.Parse()
	// Open/create the log file
	lf, err := os.OpenFile(LOGGER_PATH, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		logger.Error("Failed to open log file: %v", err)
	}
	// Init the default logger, and print the output to the console & to the file
	logger.Init("SchedulerLogger", true, false, lf)

	// Don't close the file gracefully by "defer" function, since they called once this init function finished
}
