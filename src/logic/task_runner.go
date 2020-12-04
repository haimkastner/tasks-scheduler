package logic

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/google/logger"

	"github.com/haimkastner/tasks-scheduler.git/src/model"
)

// Run task, used to run it async (by go keyword), else the app will block till application closed
func runTaskAsync(task *model.Task) {
	// Create the command to invoke
	cmd := exec.Command(task.Application, task.Args...)

	// Invoke the application
	stdout, err := cmd.Output()

	// If the application fail, log it
	if err == nil {
		logger.Info(fmt.Sprintf(`Executing task "%s" succeed`, task.Name))
	} else {
		logger.Warning(fmt.Sprintf(`Executing task "%s" return with non zero code, %v`, task.Name, err))
	}

	// Print to console the application stdout
	log.Print(string(stdout))
}

// Invoke task application
func RunTask(task *model.Task) {
	// Start task runner async
	// notice to pass the real object address and *not" just copy the popinter,
	// because tasks (and any) for iterator used the same pointer for iterating object, so the async runner can got wrong task in middle of the execution
	go runTaskAsync(&*task)
}
