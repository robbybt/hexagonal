package cart

import (
	"context"
)

type CartRepository interface {
	GetCartDetailByUserCartStatuses(ctx context.Context, userIDs []int64, statuses []int, isAllowBundle bool) ([]Cart, error)
}
