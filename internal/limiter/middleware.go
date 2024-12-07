package limiter

import "net/http"

func Middleware(ratelimiter *RateLimiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			key := r.RemoteAddr
			token := r.Header.Get("API_KEY")

			if token != "" {
				key = "token" + token
				allowed, err := ratelimiter.Allow(key, ratelimiter.rateLimitToken)
				if err != nil || !allowed {
					http.Error(w, "you have reached the maximum number of requests or actions allowed within a certain time frame", http.StatusTooManyRequests)
					return
				}
			} else {
				allowed, err := ratelimiter.Allow(key, ratelimiter.rateLimitIP)
				if err != nil || !allowed {
					http.Error(w, "you have reached the maximum number of requests or actions allowed within a certain time frame", http.StatusTooManyRequests)
					return
				}
			}

			next.ServeHTTP(w, r)
		})
	}
}
