# FROM  golang:1.19 as builder
FROM  quay.io/projectquay/golang:1.20 as builder
ARG TARGETOS
ARG TARGETOSARCH
RUN echo "TARGETOS=${TARGETOS}"
RUN echo "TARGETOSARCH=${TARGETOSARCH}"
WORKDIR /go/src/app
COPY . .
# RUN make build
RUN make $TARGETOS TARGETOSARCH=$TARGETOSARCH
FROM scratch
WORKDIR /
COPY --from=builder /go/src/app/kbot .
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT [ "./kbot", "start"]