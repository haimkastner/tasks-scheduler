package logic

import (
	"fmt"
	"io/ioutil"

	"github.com/fsnotify/fsnotify"
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

// Watch changes in the tasks file, and once it's detected trigger loadTasks()
// pass the tasks pointes pointer to allow override the pointer data
func watchChanges(tasksPtr **model.Tasks) {
	// Create watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		logger.Error("Failed to open file watcher ", err)
		return
	}
	defer watcher.Close()
	// Open tasks file watch
	err = watcher.Add(utils.TASKS_CONFIG_PATH)
	if err != nil {
		logger.Error(fmt.Sprintf(`Failed to open watcher on the "%s" file`, utils.TASKS_CONFIG_PATH), err)
		return
	}
	// Infinity loop, to get the watch channel event messages
	for {
		event, ok := <-watcher.Events
		if !ok {
			logger.Warning("Watcher event error, event:", event)
		}
		if event.Op&fsnotify.Write == fsnotify.Write {
			logger.Info("Tasks file change detected, triggering tasks data reload...")
			loadTasks(tasksPtr)
		}

	}
}

// Load the tasks pointer (by read & parse tasks file) with the latest tasks
func loadTasks(tasksPtr **model.Tasks) {
	// Init tasks model for loading
	tasks := model.Tasks{}

	logger.Info(fmt.Sprintf(`Loading the "%s" tasks file...`, utils.TASKS_CONFIG_PATH))
	// read the tasks.yaml file content
	taskFileContent := readTasksFile()

	// Load the tasks model with the yaml file text content
	err := yaml.Unmarshal([]byte(taskFileContent), &tasks)
	if err != nil {
		logger.Error("Failed to parse tasks file error: ", err)
		logger.Warning("Please fix the file, then save it and wait for the file change detection")
		// In case of failure, set an empty tasks
		*tasksPtr = &model.Tasks{Tasks: nil}
		return
	}

	// Update the shared pointer data pinting
	*tasksPtr = &tasks
	logger.Info("The tasks successfully loaded \n", taskFileContent)
}

// Get the tasks to run in schedule,
//
// The tasks stored in a tasks.yaml file in the root directory.
//
// Returns a tasks object pointer
func GetTasks() **model.Tasks {

	// Create pointer of tasks pointer, to allow override & switch the pointing data, once the tasks file will change.
	var tasksPtr **model.Tasks

	// Create tasks pointer, this pointer will keep the tasks
	var tasks *model.Tasks

	// Keep pointing to the tasks pointer, so even stack copy will hold the same tasks pointer
	tasksPtr = &tasks

	// Load the tasks
	loadTasks(tasksPtr)

	// Watch changes, and once change detected reload tasks and update the data pointing
	go watchChanges(tasksPtr)

	return tasksPtr
}
