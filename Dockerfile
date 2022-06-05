FROM golang:1.18

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./
copy ./ ./

RUN go test -v --tags=unit ./...
RUN go build -o main ./cmd/main.go

EXPOSE 8069

CMD ["./main"]