package limiter_test

import (
	"testing"
	"time"

	"github.com/claytonssmint/desafio-rate-limiter-go/internal/limiter"
)

func TestRateLimiter(t *testing.T) {
	mockStorage := limiter.NewInMemoryStorage()
	rateLimiter := limiter.NewRateLimiter(mockStorage, 5, 10, time.Second*10)

	for i := 0; i < 5; i++ {
		allowed, _ := rateLimiter.Allow("ip:192.168.1.1", 5)
		if !allowed {
			t.Errorf("Request %d should be allowed", i+1)
		}
	}

	allowed, _ := rateLimiter.Allow("ip:192.168.1.1", 5)
	if allowed {
		t.Error("Request should have been blocked")
	}
}
