# tasks-scheduler

[![CI CD Status](https://github.com/haimkastner/tasks-scheduler/workflows/Tasks%Scheduler%20CI%20CD/badge.svg?branch=master)](https://github.com/haimkastner/tasks-scheduler/actions)

A lightweight Go task schedule to run any application in schedule. 

Run by `go run cmd/main.go`

Build by `go build cmd/main.go`

Tasks configured by the `tasks.yaml` file [see example](./tasks.yaml), the file should be in the execution directory.

Logs will saved to the `tasks_scheduler.log`  file in execution directory.

Enjoy!
