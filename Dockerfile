FROM golang:latest as builder

RUN go version

COPY . /go/src/github.com/adeptmind/adept-go-postgres-api-boilerplate/
WORKDIR /go/src/github.com/adeptmind/adept-go-postgres-api-boilerplate/

RUN go get

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app .

FROM alpine:3.7
COPY --from=builder /go/src/github.com/adeptmind/adept-go-postgres-api-boilerplate/app ./app
CMD ["./app"]