.PHONY: init test

init:
	rm -f go.mod go.sum
	go mod init && go mod tidy

test:
	docker-compose -f docker-compose.yml run bazel