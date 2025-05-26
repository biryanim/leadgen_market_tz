FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git make

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /bin/app ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /bin/app /app/app
COPY --from=builder /app/local.env /app/local.env
COPY --from=builder /app/internal/config/env /app/internal/config/env

RUN apk add --no-cache tzdata wget

RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

EXPOSE 8080


CMD ["./app"]