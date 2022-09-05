FROM golang:1.15

WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go build -o ./app/api ./cmd/main.go

FROM alpine:latest

WORKDIR /app
COPY --from=0 /app .

CMD ["./app/api"]