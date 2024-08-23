package middleware

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var (
	limiter = rate.NewLimiter(rate.Limit(10), 30) // 10 requests per second, burst of 30
	ipMap   = make(map[string]*rate.Limiter)
	mu      sync.Mutex
)

func RateLimitMiddleware() gin.HandlerFunc {
	// Get rate limit configuration from environment variables
	ratePerSecond, _ := strconv.ParseFloat(os.Getenv("RATE_LIMIT_PER_SECOND"), 64)
	if ratePerSecond == 0 {
		ratePerSecond = 10 // Default to 10 if not set
	}

	burst, _ := strconv.Atoi(os.Getenv("RATE_LIMIT_BURST"))
	if burst == 0 {
		burst = 30 // Default to 30 if not set
	}

	return func(c *gin.Context) {
		ip := c.ClientIP()
		mu.Lock()
		if _, found := ipMap[ip]; !found {
			ipMap[ip] = rate.NewLimiter(rate.Limit(ratePerSecond), burst)
		}
		mu.Unlock()

		if !ipMap[ip].Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
			c.Abort()
			return
		}
		c.Next()
	}
}
