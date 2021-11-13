.PHONY: build clean start restart stop help
SERVER = ego-layout
CONFIG = .env.develop
PORT = 10244
ENV = production

all: start

build:
	@echo "building..."
	@go build -o $(SERVER) main.go
	@echo "build success"

clean:
	@echo "cleaning..."
	@rm -rf $(SERVER)
	@go clean -i .
	@echo "clean success"

start:
	@make build
	@echo "starting..."
	@nohup ./$(SERVER) -config $(CONFIG) -port $(PORT) -env $(ENV) > nohup.log 2>&1 &
	@echo "start success"

restart:
	@make build
	@echo "restarting..."
	@ps aux | grep "$(SERVER)" | grep -v grep | awk '{print $$2}' | xargs -I {} kill -1 {}
	@echo "restart success"

stop:
	@echo "stopping..."
	@ps aux | grep "$(SERVER)" | grep -v grep | awk '{print $$2}' | xargs -I {} kill -9 {}
	@echo "stop success"

help:
	@echo "make build: compile packages and dependencies"
	@echo "make clean: remove object files and cached files"
	@echo "make start: make build and start service"
	@echo "make restart: make build and then grace restart service"
	@echo "make stop: grace stop service"
