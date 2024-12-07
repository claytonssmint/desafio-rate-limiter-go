FROM golang:1.23 as build
WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o desafio-rate-limiter-go ./cmd/app/main.go

FROM scratch
WORKDIR /app
COPY --from=build /app/desafio-rate-limiter-go .
COPY --from=build /app/.env .
ENTRYPOINT ["/app/desafio-rate-limiter-go"]
