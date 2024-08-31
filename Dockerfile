FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod vendor

COPY . .

RUN GOARCH=amd64 GOOS=linux go build -o /app/myapp ./cmd
RUN chmod +x /app/myapp

CMD ["./myapp"]