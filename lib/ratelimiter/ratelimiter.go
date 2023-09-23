package ratelimiter

import (
	"context"
)

type source string

const (
	ATCSource source = "atc_source"
)

type LimiterCheckParam struct {
	Mode     int
	Key      string
	Duration int
}

type LimiterCheckResult struct {
	Status  bool
	Counter int
}

// LimiterCheckHandler will act like handler for repository
type LimiterCheckHandler func(ctx context.Context, param LimiterCheckParam) (LimiterCheckResult, error)

// CheckRequest will check the rate limit from source and will be handled by LimiterCheckHandler
func CheckRequest(ctx context.Context, s source, limiterCheck LimiterCheckHandler) bool {
	return false
}
