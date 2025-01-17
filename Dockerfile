FROM golang:latest AS builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .

RUN go build -o server cmd/app/main.go

FROM golang:latest

WORKDIR /root/

COPY --from=builder /app/server .
COPY --from=builder /app/text /root/text
COPY --from=builder /app/.env .
EXPOSE 3432

CMD ["./server"]