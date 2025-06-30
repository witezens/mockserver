FROM golang:1.22 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o mock-server .

FROM alpine:3.18

WORKDIR /app

# additional
RUN apk add --no-cache wget

# Copiar binario y datos
COPY --from=builder /app/mock-server .
COPY mock-data ./mock-data

# Exponer el puerto de servicio
EXPOSE 3000

ENTRYPOINT ["./mock-server"]

