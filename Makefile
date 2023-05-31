APP=$(shell basename $(shell git remote get-url origin) |cut -d '.' -f1)
REGISTRY ?=vanelin

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
	docker build . -t ${REGISTRY}/${APP}:${VERSION}-${TARGETOS}-${TARGETOSARCH} --build-arg TARGETOS=${TARGETOS} --build-arg TARGETOSARCH=${TARGETOSARCH} --no-cache

push:
	docker push ${REGISTRY}/${APP}:${VERSION}-${TARGETOS}-${TARGETOSARCH}

clean:
	rm -rf kbot
	docker rmi ${REGISTRY}/${APP}:${VERSION}-${TARGETOS}-${TARGETOSARCH}

# linux: TARGETOS=linux
# linux: build image push clean

linux: # Build for linucx, by default this made for arm64
	${MAKE} build TARGETOS=linux

windows: # Build for windows, by default this made for arm64
	${MAKE} build TARGETOS=windows

macos: # Build for macos, by default this made for arm64
	${MAKE} build TARGETOS=darwin