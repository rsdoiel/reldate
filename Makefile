#
# Simple Makefile
#

build: reldate.go cmds/reldate/reldate.go
	go build -o bin/reldate cmds/reldate/reldate.go 

clean: reldate
	rm bin/reldate

install: reldate.go
	go install cmds/reldate/reldate.go

