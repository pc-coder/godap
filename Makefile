
WORKING_DIR := $(shell pwd)

dockerize:
	docker build -t godap:latest .

start-dockerized-server:
	docker run --rm \
	-p 389:389 \
	--mount type=bind,source=$(WORKING_DIR)/config,target=/config godap:latest