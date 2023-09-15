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

type LimiterCheckHandler func(ctx context.Context, param LimiterCheckParam) (LimiterCheckResult, error)

func CheckRequest(ctx context.Context, s source, limiterCheck LimiterCheckHandler) bool {
	return false
}
