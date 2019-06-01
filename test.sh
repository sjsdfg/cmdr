#!/usr/bin/env bash

wget() {
	# -ncnvvqb: -nc -nv -v -q -b
	go run ./examples/wget-demo/main.go --config=wget.config -cqv -t3 -odemoout.txt -i wget-list.lst --config=wget.config --retry-on-http-error=int,bug,memories -ncnvvqb $*
	
	# environment variable bind test:
	APP_CONFIG=wget.conf go run ../wget-demo/main.go ~~debug $*
}

build-short() {
	PKG_SRC=./examples/short/main.go APPNAME=short ./build.sh
}

build-demo() {
	PKG_SRC=./examples/demo/main.go APPNAME=demo ./build.sh
}

build-fluent() {
	PKG_SRC=./examples/fluent/main.go APPNAME=fluent ./build.sh
}

build-all() {
	PKG_SRC=./examples/short/main.go APPNAME=short ./build.sh
	PKG_SRC=./examples/demo/main.go APPNAME=demo ./build.sh
	PKG_SRC=./examples/wget-demo/main.go APPNAME=wget-demo ./build.sh
	PKG_SRC=./examples/fluent/main.go APPNAME=fluent ./build.sh
}

build-all-linux() {
	PKG_SRC=./examples/short/main.go APPNAME=short ./build.sh linux
	PKG_SRC=./examples/demo/main.go APPNAME=demo ./build.sh linux
	PKG_SRC=./examples/wget-demo/main.go APPNAME=wget-demo ./build.sh linux
	PKG_SRC=./examples/fluent/main.go APPNAME=fluent ./build.sh linux
}

build-ci() {
  go mod download
	PKG_SRC=./examples/short/main.go APPNAME=short ./build.sh all
	PKG_SRC=./examples/demo/main.go APPNAME=demo ./build.sh all
	PKG_SRC=./examples/wget-demo/main.go APPNAME=wget-demo ./build.sh all
	PKG_SRC=./examples/fluent/main.go APPNAME=fluent ./build.sh all
	ls -la ./bin/
	for f in bin/*; do gzip $f; done 
	ls -la ./bin/
}

run-wget() {
	go run ./examples/wget-demo/main.go $*
}

run-wget-demo() {
	go run ./examples/wget-demo/main.go $*
}

run-demo() {
	go run ./examples/demo/main.go $*
}

run-fluent() {
	go run ./examples/fluent/main.go $*
}

fmt() {
	gofmt -l -w -s .
}

lint() {
  golint ./...
}

gotest() {
  go test ./...
}

test() {
  go test ./...
}

gocov() {
  go test -race -covermode=atomic -coverprofile cover.out && \
  go tool cover -html=cover.out -o cover.html && \
  open cover.html
}

gocov-codecov() {
  # https://codecov.io/gh/hedzr/cmdr
  go test -race -coverprofile=coverage.txt -covermode=atomic
  bash <(curl -s https://codecov.io/bash) -t $CODECOV_TOKEN
}

gocov-codecov-open() {
  open https://codecov.io/gh/hedzr/cmdr
}


[[ $# -eq 0 ]] && {
	run-demo
} || {
	cmd=$1 && shift
	case $cmd in
	*) $cmd "$@" ;;
	esac
}

