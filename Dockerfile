FROM golang:1.18

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./
copy ./ ./

RUN go build -o main ./cmd/web/main.go

EXPOSE 8069

CMD ["./main"]