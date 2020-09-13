.PHONY: all build run clean shell cleangen
all: build run clean

IMAGE=fi/slide:v1

build: *
	docker build -f Dockerfile -t ${IMAGE} ..
run: cleangen
	@mkdir -p gen
	docker run --rm -v ${CURDIR}/gen:/gen -u $(shell id -u):$(shell id -g) ${IMAGE}
shell:
	docker run --rm -it ${IMAGE} bash
clean:
	docker rmi ${IMAGE}
cleangen:
	@rm -rf ${CURDIR}/gen/