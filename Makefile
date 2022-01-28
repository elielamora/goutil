.PHONY: make-image
make-image:
	docker build -t github.com/elielamora/goutil:latest

.PHONY: test
test:
	docker build -t github.com/elielamora/goutil:latest .
	docker run -v ${PWD}:/go/testdir -t github.com/elielamora/goutil:latest