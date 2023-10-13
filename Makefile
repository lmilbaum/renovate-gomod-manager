.PHONY: build
build:
	docker build -t renovate-gomod-manager .

.PHONY: run
run:
	docker run -v ${WORKSPACE}:/workspace -it renovate-gomod-manager

.PHONY: go-run
go-run:
	go run main.go /workspace