package main

import (
	"fmt"
	"net/http"

	"github.com/claytonssmint/desafio-rate-limiter-go/config"
	"github.com/claytonssmint/desafio-rate-limiter-go/internal/limiter"
)

func main() {
	cfg := config.LoadConfig()

	storage := limiter.NewRedisStorage(cfg.RedisAddr, cfg.RedisPass)
	rateLimiter := limiter.NewRateLimiter(storage, cfg.RateLimitIP, cfg.RateLimitToken, cfg.BlockDuration)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Rate Limited API"))
	})

	handler := limiter.Middleware(rateLimiter)(mux)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", handler)
}
