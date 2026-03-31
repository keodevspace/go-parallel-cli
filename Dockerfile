# Alteramos de 1.21 para 1.25
FROM golang:1.25-alpine AS builder

WORKDIR /app

# O resto permanece igual
COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-parallel-cli cmd/cli/main.go

# Estágio Final
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go-parallel-cli /go-parallel-cli
ENTRYPOINT ["/go-parallel-cli"]