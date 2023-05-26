# kbot

DevOps application from scratch.
Simple CLI to send on-demand messages on behalf of a Telegram bot.

## Config / flags

URL to Telegram bot: [t.me/ivanvoloboyev_bot](https://t.me/ivanvoloboyev_bot) 

tested Go version = 1.19.7 

## Basic Installation
Steps on local env:
```bash
# password-protected SSH key
$ git clone git@github.com:vanelin/kbot.git 

$ cd kbot

# download and install dependency
$ go get

# type or copy-past telegram token for t.me/ivanvoloboyev_bot
$ read -s TELE_TOKEN

$ export TELE_TOKEN

# build kbot app
$ go build -ldflags "-X="github.com/vanelin/kbot/cmd.appVersion=v1.0.2

# start app
$ ./kbot start
```
## Usage
After `./kbot start` go to Telegram bot and type **Srart**

Available commands:

1. `/start hello`
2. `/start ping`

## Makefile allows you to build code on different platforms. 

To `make push` you need to authorize google:

```bash
#install Google Cloud Code extantions to vscode
$ gcloud auth login
$ gcloud config set project minikube-385711
$ gcloud auth configure-docker
```

With this Makefile, you can specify the `TARGETOS`, `TARGETOSARCH` and `REGESTRY` variables when running the make command. 

For example, to build the code for Linux on an AMD64/ARM64 architecture, and push to Google container registry, you would run:
```bash
$ make linux TARGETOSARCH=amd64 REGESTRY=HOSTNAME/PROJECT-ID
$ make linux TARGETOSARCH=arm64 REGESTRY=HOSTNAME/PROJECT-ID

```
For example, to build the code for MacOS on an ARM64/ARM64 architecture, and push to Docker Hub container registry, you would run:
```bash
$ make macos TARGETOSARCH=amd64 REGESTRY=yourhubusername
$ make macos TARGETOSARCH=arm64 REGESTRY=yourhubusername
```

For example, to build the code for Windows on an AMD64/ARM64 architecture, you would run:
```bash
$ make windows TARGETOSARCH=amd64
$ make windows TARGETOSARCH=arm64
```

> The full list is based on the result of the command `go tool dist list`. The "32-bit/64-bit" information is based on https://golang.org/doc/install/source.

## Support Grid:

|                   | `darwin` |  `linux` | `windows` |                   |
| ----------------: | :------: |  :-----: | :-------: | :---------------- |
| **`amd64`**       |  ✅      | ✅       | ✅         | **`amd64`**      |
| **`arm`**         |          | ✅       | ✅         | **`arm`**        |
| **`arm64`**       | ✅       | ✅       | ✅         | **`arm64`**      |
|  |  **`darwin`** | **`linux`** | **`windows`** |  |

# Kbot Chart
A Helm chart for Kbot, a declarative, GitOps continuous delivery tool for Kubernetes.

Source code can be found here:

https://github.com/vanelin/kbot/helm

## Prerequisites
- Kubernetes: `>=1.26.0-0`
- Helm `v3.0.0+`

## Installing the Chart

To install the chart with the release name `my-release`:

```console
$ helm repo add kbot https://github.com/vanelin/kbot/helm
"kbot" has been added to your repositories

$ helm install my-release vanelin/kbot --set secret.key="your_token"

NAME: my-release
SECRET.KEY: type token for bot kbot
...
```
## General parameters

| Key | Type | Default |
|-----|------|---------|
| image.repository | string | `"gcr.io/minikube-385711"` |
| image.tag | string | `"v1.0.5-a11547b"` |
| image.tag  | string | `"arm64"` |
| image.tag  | string | `"arm64"` | 
| secret.name  | string | `"kbot"` |
| secret.env  | string | `"TELE_TOKEN"` |
| secret.key  | string | `"your_token"` |
| securityContext.privileged | bool | `true` |


