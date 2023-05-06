APP=$(shell basename $(shell git remote get-url origin))
REGESTRY=gcr.io/minikube-385711 #vanelin

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

linux: TARGETOS = linux
linux: build image push clean

windows: TARGETOS = windows
windows: build image push clean

macos: TARGETOS = darwin
macos: build image push clean