package model

// The available scheduler types
type SchedulerType string

const (
	// Hourly schedule
	Hourly SchedulerType = "hourly"
	// Daily schedule
	Daily SchedulerType = "daily"
	// Weekly schedule
	Weekly SchedulerType = "weekly"
)

// A task object, contained the application to run and the schedule trigger properties
type Task struct {
	// The task name (optional)
	Name string
	// The application to run (for example C:\\Program Files\\PuTTY\\putty.exe in Windows or echo in Linux)
	Application string
	// The arguments to pass to the application
	Args []string
	// The scheduler trigger type
	Scheduler SchedulerType
	// The day (zero is Sunday) in week, relevent to SchedulerType.Weekly only
	Day int
	// The hour (0-23) in the day, relevent to SchedulerType.Weekly and SchedulerType.Daily only
	Hour int
	// The minute (0-59) in the hour, relevent to all scheduler types
	Minute int
}

// Represents an collection of tasks to trigger by scheduler
type Tasks struct {
	// The tasks collection
	// (hold as pointers array for performance and to allow pass it within a for loop to a goroutine)
	Tasks []*Task
}
