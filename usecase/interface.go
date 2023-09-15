package usecase

import (
	"context"
)

type UsecaseInterface interface {
	DoAtc(ctx context.Context, req RequestFrontEndATC) (responseATC, error)
}
