FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /skyshi

COPY . .

RUN go mod tidy

RUN go build -o binary

ENTRYPOINT ["/skyshi/binary"]