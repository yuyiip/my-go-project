SHELL := /bin/bash
BASEDIR = $(shell pwd)

# build with version infos
versionDir = "my-go-project/pkg/version"
gitTag = $(shell if [ "`git describe --tags --abbrev=0 2>/dev/null`" != "" ];then git describe --tags --abbrev=0; else git log --pretty=format:'%h' -n 1; fi)
buildDate = $(shell TZ=Asia/Singapore date +%FT%T%z)
gitCommit = $(shell git log --pretty=format:'%H' -n 1)
gitTreeState = $(shell if git status|grep -q 'clean';then echo clean; else echo dirty; fi)

ldflags="-w -X ${versionDir}.gitTag=${gitTag} -X ${versionDir}.buildDate=${buildDate} -X ${versionDir}.gitCommit=${gitCommit} -X ${versionDir}.gitTreeState=${gitTreeState}"


all: gotool
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -ldflags ${ldflags}
clean:
	rm -f my-go-project
	find . -name "[._]*.s[a-w][a-z]" | xargs -i rm -f {}
gotool:
	go fmt .
	go vet . | grep -v vendor;true
help:
	@echo "make - compile the source code"
	@echo "make clean - remove binary file and vim swp files"
	@echo "make gotool - run go tool 'fmt' and  'vet'"

.PHONY: clean gotool help