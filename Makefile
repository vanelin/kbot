APP=$(shell basename $(shell git remote get-url origin) |cut -d '.' -f1)
REGESTRY=gcr.io/minikube-385711
# REGESTRY=vanelin
VERSION=$(shell git describe --tags --abbrev=0)-$(shell git rev-parse --short HEAD)
TARGETOS=linux
TARGETOSARCH=arm64

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
	docker build . -t ${REGESTRY}/${APP}:${VERSION}-${TARGETOSARCH} --build-arg TARGETOS=${TARGETOS} --build-arg TARGETOSARCH=${TARGETOSARCH} --no-cache

push:
	docker push ${REGESTRY}/${APP}:${VERSION}-${TARGETOSARCH}

clean:
	rm -rf kbot
	docker rmi ${REGESTRY}/${APP}:${VERSION}-${TARGETOSARCH}

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