.PHONY: build
build:
	docker build -t gomod-manager .

.PHONY: run
run:
	docker run -v /Users/lmilbaum/repos/github.com/lmilbaum/colima:/workspace -it gomod-manager 