FROM golang:1.13-alpine as builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/src/
COPY ./ ./

RUN go mod download

RUN go build -o app

FROM alpine:3.11.3
RUN apk add tzdata

COPY --from=builder /go/src/app /go/src/app

CMD ["/go/src/app"]