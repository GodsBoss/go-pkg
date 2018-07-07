build:
	go build ./...

test:
	go test -cover ./...

.PHONY: \
	build \
	test
