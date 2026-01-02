package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/didip/tollbooth/v8"
	"github.com/didip/tollbooth/v8/limiter"
)

type LimitMessage struct {
	Status string
	Body   string
}

func NewRateLimiter() *limiter.Limiter {
	message := LimitMessage{
		Status: "Request Failed",
		Body:   "Rate limit reached. Please try again later.",
	}

	jsonMessage, _ := json.Marshal(message)

	lmt := tollbooth.NewLimiter(5.0/3600.0, nil) // 5 requests per hour
	lmt.SetMessageContentType("application/json")
	lmt.SetMessage(string(jsonMessage))

	lmt.SetIPLookup(limiter.IPLookup{
		Name:           "RemoteAddr",
		IndexFromRight: 0,
	})

	return lmt
}

func RateLimitEndpoint(rateLimiter *limiter.Limiter, handler http.HandlerFunc) http.Handler {
	return tollbooth.HTTPMiddleware(rateLimiter)(handler)
}
