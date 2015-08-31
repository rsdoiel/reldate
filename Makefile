#
# Simple Makefile
#

build: reldate.go cmd/reldate/reldate.go
	go build -o bin/reldate cmd/reldate/reldate.go 

clean: reldate
	rm bin/reldate

install: reldate.go
	go install cmd/reldate/reldate.go

