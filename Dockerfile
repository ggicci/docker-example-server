FROM golang:alpine
COPY . $GOPATH/src/example
RUN go install example && rm -rf $GOPATH/src/example
ENTRYPOINT [ "example" ]
EXPOSE 80
