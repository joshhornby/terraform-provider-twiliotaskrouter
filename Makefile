NAME:=$(shell basename $$PWD)
ARCH:=$(shell uname -m)
REPO:=$(shell git config --get remote.origin.url | perl -ne 'm{github.com[:/](.+/[^.]+)}; print  $$1')
VERSION=0.3.0

test:
	go test ./...

build:
	go build

release: build
	rm -rf release && mkdir release
	mkdir -p build && mv $(NAME) build
	tar -zcf release/$(NAME)_$(VERSION).tgz -C build $(NAME)
