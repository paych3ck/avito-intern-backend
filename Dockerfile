FROM golang:1.22-alpine

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o main ./cmd/app

EXPOSE 8080

ENTRYPOINT ["./main"]
