FROM golang:alpine AS build-env
WORKDIR /usr/local/go/src/github.com/prateekpandey14/electron/electron
COPY . /usr/local/go/src/github.com/prateekpandey14/electron/electron
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
RUN go get ./...
RUN go build -o build/electron ./electron


FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=build-env /usr/local/go/src/github.com/prateekpandey14/electron/electron/build/electron /bin/electron
CMD ["electron", "up"]
