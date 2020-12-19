# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build

# Entry points
ENTRY_POINT=cmd/main.go

# Output names
BINARY_NAME_LINUX_X86=tasks_scheduler_linux_x86
BINARY_NAME_LINUX_ARM=tasks_scheduler_linux_arm
BINARY_NAME_WIN=tasks_scheduler.exe

all: test, build
build-all: build-linux-x86 build-linux-arm build-win
build-linux-x86:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME_LINUX_X86) -v $(ENTRY_POINT)
build-linux-arm:
	GOOS=linux GOARCH=arm $(GOBUILD) -o $(BINARY_NAME_LINUX_ARM) -v $(ENTRY_POINT)
build-win:
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME_WIN) -v $(ENTRY_POINT)
