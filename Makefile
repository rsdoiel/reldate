#
# Simple Makefile
#

build:
	go build -o bin/reldate cmds/reldate/reldate.go 

clean:
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi

install:
	env GOBIN=$HOME/bin go install cmds/reldate/reldate.go

release:
	./mk-release.sh

