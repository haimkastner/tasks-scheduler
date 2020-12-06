# tasks-scheduler

A lightweight Golang task-schedule to run any application in schedule. 

Run by `go run cmd/main.go`

Build by `go build cmd/main.go`

Tasks configured by the `tasks.yaml` file [see example](./tasks.yaml), the file should be in the execution directory, tasks reload on a file change detection. 

Logs will saved to the `tasks_scheduler.log`  file in execution directory.

Enjoy!
