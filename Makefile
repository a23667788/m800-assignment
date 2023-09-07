
PROJECTNAME := $(shell basename "$(PWD)")

#project related variables
PRJBASE := $(shell pwd)
PRJBIN  := $(PRJBASE)/bin
GOFILES := ./cmd/...
DB_ID   := $(shell docker ps -a -q --filter="name=m800-mongo")
NGROK_ID   := $(shell docker ps -a -q --filter="name=m800-ngrok")

.PHONY: dev-setup
dev-setup:
	@docker pull mongo:4.4
	@docker pull ngrok/ngrok
	@go get -u github.com/gin-gonic/gin
	@go get -u github.com/spf13/viper
	@go get -u go.mongodb.org/mongo-driver
	@go get -u go.mongodb.org/mongo-driver/mongo
	@go get -u github.com/line/line-bot-sdk-go
	@go get -u github.com/line/line-bot-sdk-go/v7/linebot
	@go get -u github.com/spf13/cobra

.PHONY: db-run
db-run:
	@docker run -d -p 27017:27017 --name m800-mongo -d mongo:4.4

.PHONY: db-stop
db-stop:
	@docker stop $(DB_ID)
	@docker rm $(DB_ID)

.PHONY: ngrok-run
ngrok-run:
	# @./scripts/ngrok.exe http 8087
	@./scripts/ngrok http 8087

.PHONY: ngrok-docker-run
ngrok-docker-run:
	@docker run --name m800-ngrok -it -e NGROK_AUTHTOKEN=2Uq12WBLE389Wl4d7rcHJ2cTtzv_o9ad8uiY3yi3r7XKq1Zu ngrok/ngrok http 8087

.PHONY: ngrok-stop
ngrok-docker-stop:
	@docker stop $(NGROK_ID)
	@docker rm $(NGROK_ID)

.PHONY: clean
build: clean
	@gofmt -s -w .
	@go build -o $(PRJBIN)/$(PROJECTNAME) $(GOFILES)

.PHONY: clean
clean:
	@go clean
	@rm -rf  ${PRJBIN}/

.PHONY: run
run:
	@./bin/$(PROJECTNAME) start --port 8087 --config $(PRJBASE)/configs/config.yaml
