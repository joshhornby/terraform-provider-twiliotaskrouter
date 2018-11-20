NAME:=$(shell basename $$PWD)
ARCH:=$(shell uname -m)
REPO:=$(shell git config --get remote.origin.url | perl -ne 'm{github.com[:/](.+/[^.]+)}; print  $$1')
VERSION=0.0.2

build:
	go build
