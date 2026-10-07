package ingest

import (
	"context"
	"errors"
)

// ErrNotImplemented marks scaffolded behaviour that has no implementation yet.
var ErrNotImplemented = errors.New("not implemented")

// Apply updates the recommendation state for one completed order.
//
// TODO(RFC §5.2): ZINCRBY user:{userId}:skus, hot:global and the pairwise
// cooccur:{sku} sets.
func Apply(_ context.Context, _ OrderCompleted) error {
	return ErrNotImplemented
}
