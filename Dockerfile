FROM golang:1.22-alpine3.19 as gobuilder
LABEL authors="sim"

COPY  / /build/
WORKDIR /build
RUN go build -C cmd/goserver -ldflags "-w -s" -o goserver .


FROM alpine:3.19
WORKDIR /opt/goserver

COPY --from=gobuilder --chown=docker:docker /build/cmd/goserver/goserver goserver
COPY templates templates

ENTRYPOINT ["./goserver", "-templates", "templates", "-host", "0.0.0.0", "-port", "8080"]
