.PHONY: api-build job-build api-start job-start api-stop job-stop api-restart job-restart clean help
SERVER = ego-layout
CONFIG = .env.development
PORT = 10244
ENV = production
APIUniqueId = $(SERVER)-$(PORT)-api
JobUniqueId = $(SERVER)-$(PORT)-job

api-build:
	@echo "building api..."
	@go build -o $(APIUniqueId) cmd/api/main.go
	@echo "build success"

job-build:
	@echo "building job..."
	@go build -o $(JobUniqueId) cmd/job/main.go
	@echo "build success"

api-start:
	@make api-build
	@echo "starting api..."
	@nohup ./$(APIUniqueId) -config $(CONFIG) -port $(PORT) -env $(ENV) >> .nohup.log 2>&1 &
	@echo "start success"

job-start:
	@make job-build
	@echo "starting job..."
	@nohup ./$(JobUniqueId) -config $(CONFIG) -port $(PORT) -env $(ENV) >> .nohup.log 2>&1 &
	@echo "start success"

api-stop:
	@echo "stopping api..."
	@ps aux | grep "$(APIUniqueId)" | grep -v grep | awk '{print $$2}' | xargs -I {} kill -15 {}
	@echo "stop success"

job-stop:
	@echo "stopping job..."
	@ps aux | grep "$(JobUniqueId)" | grep -v grep | awk '{print $$2}' | xargs -I {} kill -15 {}
	@echo "stop success"

api-restart:
	@make api-build
	@echo "restarting api..."
	@ps aux | grep "$(APIUniqueId)" | grep -v grep | awk '{print $$2}' | xargs -I {} kill -1 {}
	@echo "restart success"

job-restart:
	@make job-stop
	@make job-start

clean:
	@echo "cleaning..."
	@rm -rf $(APIUniqueId)
	@rm -rf $(JobUniqueId)
	@go clean -i .
	@echo "clean success"

help:
	@echo "make clean: remove object files and cached files"
	@echo "make api-build: compile api packages and dependencies"
	@echo "make api-start: make build and start api service"
	@echo "make api-restart: make build and then grace restart api service"
	@echo "make api-stop: grace stop api service"
	@echo "make job-build: compile job packages and dependencies"
	@echo "make job-start: make build and start job service"
	@echo "make job-restart: make build and then grace restart job service"
	@echo "make job-stop: grace stop job service"