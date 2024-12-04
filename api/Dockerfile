FROM golang:1.23.2-alpine AS builder

ENV GO111MODULE=on
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o api-server cmd/main.go

FROM alpine:3.18

RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /app/api-server .

EXPOSE 8080

CMD ["./api-server"]
