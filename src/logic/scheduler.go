package logic

import (
	"fmt"
	"time"

	"github.com/google/logger"

	"github.com/haimkastner/tasks-scheduler.git/src/model"
	"github.com/haimkastner/tasks-scheduler.git/src/utils"
)

// Hold the last minute, so every changed minute the scheduler task triggers check will start
var lastMinutes int = time.Now().Minute()

// On every activation "tik" (each minute) check tasks trigger if should triggered
func onActivationTik(tasks *model.Tasks) {
	// Get the current time
	now := time.Now()
	// For each task, check the trigger
	for _, task := range tasks.Tasks {
		// Assume there is not trigger match.
		isTaskShouldTrigger := false
		// Check if it's a trigger time, accourding to the task scheduler type
		switch task.Scheduler {
		case model.Hourly:
			isTaskShouldTrigger = task.Minute == now.Minute()
		case model.Daily:
			isTaskShouldTrigger = task.Minute == now.Minute() && task.Hour == now.Hour()
		case model.Weekly:
			isTaskShouldTrigger = task.Minute == now.Minute() && task.Hour == now.Hour() && task.Day == int(now.Weekday())
		}

		// If task should trigger, do it.
		if isTaskShouldTrigger {
			logger.Info(fmt.Sprintf(`The task "%s" trigger scheduled to now`, task.Name))
			RunTask(&task)
		}
	}
}

// Detect if current minute change
//
// Return true if minute did changed
func isMinuteChanged() bool {
	// Get current time
	currentMinute := time.Now().Minute()
	// Check if the minute did changed
	isMinuteChanged := currentMinute != lastMinutes
	// Keep current time for next check
	lastMinutes = currentMinute
	// Return results
	return isMinuteChanged
}

// Active scheduler endless loop
func ActiveSchedule(tasks *model.Tasks) {

	logger.Info("The task-scheduler is running")

	for {
		// Wait a while (don't kil CPU by doing "busy wait")
		time.Sleep(utils.SCHEDULER_ACTIVATION_SEC * time.Second)
		// Run the activation tik only once in a minute
		if isMinuteChanged() {
			logger.Info("Scheduler activation tik")
			onActivationTik(tasks)
		}
	}
}
