package main

import (
	"github.com/google/logger"

	"github.com/haimkastner/tasks-scheduler.git/src/logic"
)

func main() {

	logger.Info("Starting scheduler app...")

	logger.Info("Get tasks...")
	tasksPtr := logic.GetTasks()

	logger.Info("Active schedule...")
	logic.ActiveSchedule(tasksPtr)
}
