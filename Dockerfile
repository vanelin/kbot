FROM  golang:1.19 as builder
ARG TARGETOS=linux
ARG TARGETOSARCH=arm64
WORKDIR /go/src/app
COPY . .
# RUN make build

RUN echo "I am running on $TARGETOS, building for $TARGETOSARCH" > /log
RUN make build TARGETOS=$TARGETOS TARGETOSARCH=$TARGETOSARCH

FROM scratch
WORKDIR /
COPY --from=builder /go/src/app/kbot .
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT [ "./kbot" ]