#
# Simple Makefile
#

build:
	go build -o bin/reldate cmds/reldate/reldate.go 

clean:
	if [ -d bin ]; then /bin/rm -fR bin; fi
	if [ -d dist ]; then /bin/rm -fR dist; fi
	if [ -f reldate-binary-release.zip ]; then /bin/rm reldate-binary-release.zip; fi


install:
	env GOBIN=$(HOME)/bin go install cmds/reldate/reldate.go

save:
	./mk-website.bash
	git commit -am "Quick save"
	git push origin master

website:
	./mk-website.bash

release:
	./mk-release.bash

publish:
	./mk-website.bash
	./publish.bash

