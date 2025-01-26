FROM golang:alpine AS builder

WORKDIR /app

COPY go.* .

RUN go mod download

COPY . .

RUN go build -o laskb-server-api cmd/main.go

# Stage 2

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/laskb-server-api .

EXPOSE 3250

CMD ["./laskb-server-api"]