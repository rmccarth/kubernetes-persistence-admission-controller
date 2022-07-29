FROM golang:latest as compiler

WORKDIR /code
COPY ./patch-server /code 
ENV CGO_ENABLED=0   \
    GOOS=linux      \
    GOARCH=amd64 
RUN go env -w GO111MODULE=auto && go build -ldflags '-extldflags "-static"' -o controller

FROM alpine:latest as runtime
COPY --from=compiler /code/controller /controller

EXPOSE 8443
ENTRYPOINT ["/controller"]