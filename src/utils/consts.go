package utils

import "time"

// The tasks yaml file path
const TASKS_CONFIG_PATH string = "./tasks.yaml"

// The logger path
const LOGGER_PATH string = "./tasks_scheduler.log"

// The idle time between tik check
const SCHEDULER_ACTIVATION_SEC time.Duration = 20
