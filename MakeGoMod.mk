# this file must use as base Makefile

ROOT_VENDOR_PATH := $(ROOT_PATH)/vendor
ROOT_GO_SUM_PATH := $(ROOT_PATH)/go.sum

modVerify:
	# in GOPATH must use GO111MODULE=on go mod init to init
	-GOPROXY="$(ENV_GO_PROXY)" GO111MODULE=on go mod verify

modDownload:
	-GOPROXY="$(ENV_GO_PROXY)" GO111MODULE=on go mod download
	-GOPROXY="$(ENV_GO_PROXY)" GO111MODULE=on go mod vendor

modTidy:
	-GOPROXY="$(ENV_GO_PROXY)" GO111MODULE=on go mod tidy

modUpdate:
	-GOPROXY="$(ENV_GO_PROXY)" GO111MODULE=on go get -u ./...

dep: modVerify modDownload
	@echo "just check depends info below"

depClean:
	@echo "just clean depends start"
	@if [ -d ${ROOT_VENDOR_PATH} ]; \
	then rm -rf vendor && echo "~> cleaned ${ROOT_VENDOR_PATH}"; \
	else echo "~> has cleaned ${ROOT_VENDOR_PATH}"; \
	fi
	@if [ -f ${ROOT_GO_SUM_PATH} ]; \
	then rm -f vendor && echo "~> cleaned ${ROOT_GO_SUM_PATH}"; \
	else echo "~> has cleaned ${ROOT_GO_SUM_PATH}"; \
	fi
	@echo "just clean depends end"

modGraphDependencies:
	GOPROXY="$(ENV_GO_PROXY)" GO111MODULE=on go mod graph

helpGoMod:
	@echo "Help: MakeGoMod.mk"
	@echo "this project use go mod, so golang version must 1.12+"
	@echo "go mod evn: GOPROXY=$(ENV_GO_PROXY)"
	@echo "~> make dep                  - check depends of project and download all, child task is: modVerify modDownload"
	@echo "~> make depClean             - clean path of ${ROOT_VENDOR_PATH} and ${ROOT_VENDOR_PATH}"
	@echo "~> make modGraphDependencies - see depends graph of this project"
	@echo "~> make modTidy              - tidy depends graph of project"
	@echo "~> make modUpdate            - update depends to new"
	@echo ""