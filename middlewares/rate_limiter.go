package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RateLimiter struct {
	tokens       chan struct{}
	maxTokens    int
	refillPeriod time.Duration
}

func NewRateLimiter(maxTokens int, refillPeriod time.Duration) *RateLimiter {
	rl := &RateLimiter{
		tokens:       make(chan struct{}, maxTokens),
		maxTokens:    maxTokens,
		refillPeriod: refillPeriod,
	}

	for i := 0; i < maxTokens; i++ {
		rl.tokens <- struct{}{}
	}

	go rl.refillTokens()

	return rl
}

func (rl *RateLimiter) refillTokens() {
	ticker := time.NewTicker(rl.refillPeriod)
	defer ticker.Stop()

	for range ticker.C {
		select {
		case rl.tokens <- struct{}{}:
		default:
		}
	}
}

func (rl *RateLimiter) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		select {
		case <-rl.tokens:
			c.Next()
		default:
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Rate limit exceeded, please try again later.",
			})
			return
		}
	}
}
