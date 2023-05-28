APP=$(shell basename $(shell git remote get-url origin) |cut -d '.' -f1)
REGISTRY ?=gcr.io
REPOSITORY ?=minikube-385711

VERSION=$(shell git describe --tags --abbrev=0)-$(shell git rev-parse --short HEAD)
TARGETOS ?=linux
TARGETOSARCH ?=arm64

format:
	gofmt -s -w ./

lint:
	golint

test:
	go test -v

get:
	go get

build: format get
	CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETOSARCH} go build -v -o kbot -ldflags "-X="github.com/vanelin/kbot/cmd.appVersion=${VERSION}

image:
	docker build . -t ${REGISTRY}/${REPOSITORY}/${APP}:${VERSION}-${TARGETOS}-${TARGETOSARCH} --build-arg TARGETOS=${TARGETOS} --build-arg TARGETOSARCH=${TARGETOSARCH} --no-cache

push:
	docker push ${REGISTRY}/${REPOSITORY}/${APP}:${VERSION}-${TARGETOS}-${TARGETOSARCH}

clean:
	rm -rf kbot
	docker rmi ${REGISTRY}/${REPOSITORY}/${APP}:${VERSION}-${TARGETOS}-${TARGETOSARCH}

linux: TARGETOS=linux
linux: build image push clean

windows:
	${MAKE} build TARGETOS=windows
	${MAKE} image TARGETOS=windows
	${MAKE} push TARGETOS=windows
	${MAKE} clean TARGETOS=windows

macos:
	make build TARGETOS=darwin
	make image TARGETOS=darwin
	make push TARGETOS=darwin
	make clean TARGETOS=darwin