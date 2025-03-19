FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .

COPY . .

RUN go build -o weather-bot cmd/bot/main.go

FROM alpine

WORKDIR /build

COPY --from=builder /build/weather-bot /build/weather-bot

CMD ["./weather-bot"]