FROM golang:alpine AS builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .

RUN go build -o server cmd/app/main.go

FROM golang:alpine

WORKDIR /root/

COPY --from=builder /app/server .
COPY --from=builder /app/text text
EXPOSE 3432
CMD ["./server"]
