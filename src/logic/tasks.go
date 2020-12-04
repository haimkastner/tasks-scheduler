package logic

import (
	"fmt"
	"io/ioutil"

	"github.com/google/logger"
	"gopkg.in/yaml.v2"

	"github.com/haimkastner/tasks-scheduler.git/src/model"
	"github.com/haimkastner/tasks-scheduler.git/src/utils"
)

// Read the tasks file full content
// Return it as string
func readTasksFile() string {
	// Read the file
	content, err := ioutil.ReadFile(utils.TASKS_CONFIG_PATH)
	if err != nil {
		logger.Fatal("Failed to read tasks file ", err)
	}
	// Convert the content to string and return
	return string(content)
}

// Get the tasks to run in schedule,
//
// The tasks stored in a tasks.yaml file in the root directory.
//
// Returns a tasks object pointer
func GetTasks() *model.Tasks {

	// Init tasks model for loading
	tasks := model.Tasks{}

	logger.Info(fmt.Sprintf(`Loading the "%s" tasks file...`, utils.TASKS_CONFIG_PATH))
	// read the tasks.yaml file content
	taskFileContent := readTasksFile()

	// Load the tasks model with the yaml file text content
	err := yaml.Unmarshal([]byte(taskFileContent), &tasks)
	if err != nil {
		logger.Fatal("Failed to parse tasks file error: ", err)
	}

	logger.Info("The tasks successfully loaded \n", taskFileContent)

	return &tasks
}
