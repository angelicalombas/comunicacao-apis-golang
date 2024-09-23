FROM golang:1.21.5-alpine

WORKDIR /go/src/app

COPY . .

RUN go mod tidy
RUN go build -o order-service

EXPOSE 8080

CMD ["./order-service"]