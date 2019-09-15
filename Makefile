.PHONY: test check clean build dist all

TOP_DIR := $(shell pwd)

# ifeq ($(FILE), $(wildcard $(FILE)))
# 	@ echo target file not found
# endif

DIST_VERSION := 1.0.0
DIST_OS := linux
DIST_ARCH := amd64

ROOT_NAME ?= golang-playground
ROOT_BUILD_PATH ?= ./build
ROOT_DIST ?= ./dist
ROOT_REPO ?= ./dist
ROOT_TEST_DIST_PATH ?= $(ROOT_DIST)/test/$(DIST_VERSION)
ROOT_TEST_OS_DIST_PATH ?= $(ROOT_DIST)/$(DIST_OS)/test/$(DIST_VERSION)
ROOT_REPO_DIST_PATH ?= $(ROOT_REPO)/$(DIST_VERSION)
ROOT_REPO_OS_DIST_PATH ?= $(ROOT_REPO)/$(DIST_OS)/release/$(DIST_VERSION)

ROOT_LOG_PATH ?= ./log
ROOT_SWAGGER_PATH ?= ./docs

SERVER_TEST_SSH_ALIASE = server-ci-test
SERVER_TEST_FOLDER = /home/sinlov/Documents/website
SERVER_REPO_SSH_ALIASE = web-gin-temp
SERVER_REPO_FOLDER = /home/ubuntu/$(ROOT_NAME)

# can use as https://goproxy.io/ https://gocenter.io https://mirrors.aliyun.com/goproxy/
INFO_GO_PROXY ?= https://mirrors.aliyun.com/goproxy/

checkEnvGOPATH:
ifndef GOPATH
	@echo Environment variable GOPATH is not set
	exit 1
endif

init:
	@echo "~> start init this project"
	@echo "-> check version"
	go version
	@echo "-> check env golang"
	go env
	@echo "~> you can use [ make help ] see more task"
	-GOPROXY="$(INFO_GO_PROXY)" GO111MODULE=on go mod vendor

checkDepends:
	# in GOPATH just use GO111MODULE=on go mod init to init
	-GOPROXY="$(INFO_GO_PROXY)" GO111MODULE=on go mod verify

tidyDepends:
	-GOPROXY="$(INFO_GO_PROXY)" GO111MODULE=on go mod tidy

dep: checkDepends
	@echo "just check depends info below"

dependenciesGraph:
	GOPROXY="$(INFO_GO_PROXY)" GO111MODULE=on go mod graph

cleanBuild:
	@if [ -d ${ROOT_BUILD_PATH} ]; then rm -rf ${ROOT_BUILD_PATH} && echo "~> cleaned ${ROOT_BUILD_PATH}"; else echo "~> has cleaned ${ROOT_BUILD_PATH}"; fi

cleanDist:
	@if [ -d ${ROOT_DIST} ]; then rm -rf ${ROOT_DIST} && echo "~> cleaned ${ROOT_DIST}"; else echo "~> has cleaned ${ROOT_DIST}"; fi

cleanLog:
	@if [ -d ${ROOT_LOG_PATH} ]; then rm -rf ${ROOT_LOG_PATH} && echo "~> cleaned ${ROOT_LOG_PATH}"; else echo "~> has cleaned ${ROOT_LOG_PATH}"; fi

clean: cleanBuild cleanLog
	@echo "~> clean finish"

checkTestDistPath:
	@if [ ! -d ${ROOT_TEST_DIST_PATH} ]; then mkdir -p ${ROOT_TEST_DIST_PATH} && echo "~> mkdir ${ROOT_TEST_DIST_PATH}"; fi

checkTestOSDistPath:
	@if [ ! -d ${ROOT_TEST_OS_DIST_PATH} ]; then mkdir -p ${ROOT_TEST_OS_DIST_PATH} && echo "~> mkdir ${ROOT_TEST_OS_DIST_PATH}"; fi

checkReleaseDistPath:
	@if [ ! -d ${ROOT_REPO_DIST_PATH} ]; then mkdir -p ${ROOT_REPO_DIST_PATH} && echo "~> mkdir ${ROOT_REPO_DIST_PATH}"; fi

checkReleaseOSDistPath:
	@if [ ! -d ${ROOT_REPO_OS_DIST_PATH} ]; then mkdir -p ${ROOT_REPO_OS_DIST_PATH} && echo "~> mkdir ${ROOT_REPO_OS_DIST_PATH}"; fi

buildMain:
	@echo "-> start build local OS"
	@go build -o build/main main.go

buildARCH:
	@echo "-> start build OS:$(DIST_OS) ARCH:$(DIST_ARCH)"
	@GOOS=$(DIST_OS) GOARCH=$(DIST_ARCH) go build -o build/main main.go

dev: buildMain
	-./build/main -c ./conf/config.yaml

runTest:  buildMain
	-./build/main -c ./conf/test/config.yaml

test: checkDepends buildMain checkTestDistPath
	mv ./build/main $(ROOT_TEST_DIST_PATH)
	cp ./conf/test/config.yaml $(ROOT_TEST_DIST_PATH)
	cp -R ./static $(ROOT_TEST_DIST_PATH)
	cp -R ./views $(ROOT_TEST_DIST_PATH)
	@echo "=> pkg at: $(ROOT_TEST_DIST_PATH)"

testOS: checkDepends buildARCH checkTestOSDistPath
	@echo "=> Test at: $(DIST_OS) ARCH as: $(DIST_ARCH)"
	mv ./build/main $(ROOT_TEST_OS_DIST_PATH)
	cp ./conf/test/config.yaml $(ROOT_TEST_OS_DIST_PATH)
	cp -R ./static $(ROOT_TEST_OS_DIST_PATH)
	cp -R ./views $(ROOT_TEST_OS_DIST_PATH)
	@echo "=> pkg at: $(ROOT_TEST_OS_DIST_PATH)"

testOSScp:
	#scp -r $(ROOT_TEST_OS_DIST_PATH) $(SERVER_TEST_SSH_ALIASE):$(SERVER_TEST_FOLDER)
	@echo "=> must check below config of set for testOSScp"

testOSTarScp:
#	scp $(ROOT_DIST)/$(DIST_OS)/test/$(ROOT_NAME)-$(DIST_OS)-$(DIST_ARCH)-$(DIST_VERSION).tar.gz $(SERVER_TEST_SSH_ALIASE):$(SERVER_TEST_FOLDER)
	@echo "=> must check below config of set for testOSTarScp"

testOSTar: testOS
	@echo "=> start tar test as os $(DIST_OS) $(DIST_ARCH)"
	tar zcvf $(ROOT_DIST)/$(DIST_OS)/test/$(ROOT_NAME)-$(DIST_OS)-$(DIST_ARCH)-$(DIST_VERSION).tar.gz $(ROOT_TEST_OS_DIST_PATH)

release: checkDepends buildMain checkReleaseDistPath
	mv ./build/main $(ROOT_REPO_DIST_PATH)
	cp ./conf/release/config.yaml $(ROOT_REPO_DIST_PATH)
	cp -R ./static $(ROOT_REPO_DIST_PATH)
	cp -R ./views $(ROOT_REPO_DIST_PATH)
	@echo "=> pkg at: $(ROOT_REPO_DIST_PATH)"

releaseOS: checkDepends buildARCH checkReleaseOSDistPath
	@echo "=> Release at: $(DIST_OS) ARCH as: $(DIST_ARCH)"
	mv ./build/main $(ROOT_REPO_OS_DIST_PATH)
	cp ./conf/release/config.yaml $(ROOT_REPO_OS_DIST_PATH)
	cp -R ./static $(ROOT_REPO_OS_DIST_PATH)
	cp -R ./views $(ROOT_REPO_OS_DIST_PATH)
	@echo "=> pkg at: $(ROOT_REPO_OS_DIST_PATH)"

releaseOSTar: releaseOS
	@echo "=> start tar release as os $(DIST_OS) $(DIST_ARCH)"
	tar zcvf $(ROOT_DIST)/$(DIST_OS)/release/$(ROOT_NAME)-$(DIST_OS)-$(DIST_ARCH)-$(DIST_VERSION).tar.gz $(ROOT_REPO_OS_DIST_PATH)

releaseOSTarScp: releaseOSTar
	scp $(ROOT_DIST)/$(DIST_OS)/release/$(ROOT_NAME)-$(DIST_OS)-$(DIST_ARCH)-$(DIST_VERSION).tar.gz $(SERVER_REPO_SSH_ALIASE):$(SERVER_REPO_FOLDER)
	@echo "=> must check below config of set for relase OS Scp"

releaseInitDocker:
	scp -r ./z-mk-docker $(SERVER_REPO_SSH_ALIASE):$(SERVER_REPO_FOLDER)
	@echo "=> must check below config of set for release OS Scp"

dockerComposeScpRelease:
	scp ./z-mk-docker/release/docker-compose.yml $(SERVER_REPO_SSH_ALIASE):$(SERVER_REPO_FOLDER)
	@echo "=> finish update docker compose"

help:
	@echo "make init - check base env of this project"
	@echo "make dep - check depends of project"
	@echo "make dependenciesGraph - see depends graph of project"
	@echo "make tidyDepends - tidy depends graph of project"
	@echo "make clean - remove binary file and log files"
	@echo ""
	@echo "-- now build name: $(ROOT_NAME) version: $(DIST_VERSION)"
	@echo "-- testOS or releaseOS will out abi as: $(DIST_OS) $(DIST_ARCH) --"
	@echo "make test - build dist at $(ROOT_TEST_DIST_PATH)"
	@echo "make testOS - build dist at $(ROOT_TEST_OS_DIST_PATH)"
	@echo "make testOSTar - build dist at $(ROOT_TEST_OS_DIST_PATH) and tar"
	@echo "make release - build dist at $(ROOT_REPO_DIST_PATH)"
	@echo "make releaseOS - build dist at $(ROOT_REPO_OS_DIST_PATH)"
	@echo "make releaseOSTar - build dist at $(ROOT_REPO_OS_DIST_PATH) and tar"
	@echo ""
	@echo "make runTest - run server use conf/test/config.yaml"
	@echo "make dev - run server use conf/config.yaml"