.PHONY: build lint clean start restart stop help
SERVER = ego-layout
CONFIG = .env.develop
PORT = 10244
ENV = production

all: start

build:
	go build -o $(SERVER) main.go

lint:
	golint ./...

clean:
	rm -rf $(SERVER)
	go clean -i .

start:
	make build
	nohup ./$(SERVER) -config $(CONFIG) -port $(PORT) -env $(ENV) > nohup.log 2>&1 &

restart:
	make build
	ps aux | grep "$(SERVER)" | grep -v grep | awk '{print $2}' | xargs -i kill -1 {}

stop:
	ps aux | grep "$(SERVER)" | grep -v grep | awk '{print $2}' | xargs -i kill {}

help:
	@echo "make build: compile packages and dependencies"
	@echo "make lint: golint ./..."
	@echo "make clean: remove object files and cached files"
	@echo "make start: make build and start service"
	@echo "make restart: grace restart service"
	@echo "make stop: grace stop service"
