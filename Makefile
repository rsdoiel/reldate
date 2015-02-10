#
# Simple Makefile
#

build: reldate.go
	go build reldate.go 

clean: reldate
	rm reldate

install: reldate.go
	go install reldate.go

