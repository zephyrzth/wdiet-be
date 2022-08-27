-include .env

# build
.PHONY: build

build:
	@echo " > Building [wdiet-be]..."
	@go build -o ./bin/
	@echo " > Finished building [wdiet-be]"

# RUN
run: build
	@echo " > Running [wdiet-be]..."
	@./bin/wdiet-be
	@echo " > Finished running [wdiet-be]"

