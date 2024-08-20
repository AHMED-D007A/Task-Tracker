all: run

build:
	@go build -o ./bin/

run: build
	@./bin/task_tracker