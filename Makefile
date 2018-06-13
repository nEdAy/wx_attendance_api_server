default:
	@echo 'Usage of make: [ build | clean | windows_build | linux_build ]'

build: 
	@go build -ldflags "-X main.VERSION=1.0.0 -X 'main.BUILD_TIME=`date`' -X 'main.GO_VERSION=`go version`' -X main.GIT_HASH=`git rev-parse HEAD`" -o ./build/wx_attendance_api_server ./

linux_build: 
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.VERSION=1.0.0 -X 'main.BUILD_TIME=`date`' -X 'main.GO_VERSION=`go version`' -X main.GIT_HASH=`git rev-parse HEAD` -s" -o ./build/wx_attendance_api_server ./

windows_build:
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-X main.VERSION=1.0.0 -X 'main.BUILD_TIME=`date`' -X 'main.GO_VERSION=`go version`' -X main.GIT_HASH=`git rev-parse HEAD` -s" -o ./build/wx_attendance_api_server.exe ./

windows_build_386:
	@CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -ldflags "-X main.VERSION=1.0.0 -X 'main.BUILD_TIME=`date`' -X 'main.GO_VERSION=`go version`' -X main.GIT_HASH=`git rev-parse HEAD` -s" -o ./build/wx_attendance_api_server.exe ./


clean: 
	@rm -f ./build/wx_attendance_api_server
	@rm -f ./build/wx_attendance_api_server.exe
	@rm -f ./build/logs/*.log

.PHONY: default build windows_build windows_build_386 linux_build clean