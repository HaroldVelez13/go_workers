# ----------- BUILD STAGE -----------
FROM golang:1.22-alpine AS builder

WORKDIR /app

# dependencias
COPY go.mod go.sum ./
RUN go mod download

# código
COPY . .

# build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd

# ----------- RUNTIME STAGE -----------
FROM alpine:latest

WORKDIR /root/

# binario
COPY --from=builder /app/app .

# puerto (opcional)
EXPOSE 8082

# run
CMD ["./app"]