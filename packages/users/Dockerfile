# Imagem de build
FROM golang:latest AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Imagem final
FROM scratch

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]
