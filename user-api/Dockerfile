FROM golang:1.21.5-alpine

WORKDIR /go/src/app

COPY . .

RUN go mod tidy
RUN go build -o user-service

EXPOSE 8081

CMD ["./user-service"]