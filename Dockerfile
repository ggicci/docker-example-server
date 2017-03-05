FROM golang:1.7.4-alpine
COPY . $GOPATH/src/example
RUN go install example && rm -rf $GOPATH/src/example
ENTRYPOINT [ "example" ]
EXPOSE 80
