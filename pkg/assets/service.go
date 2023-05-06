package assets

import (
	"context"

	"github.com/qwertyqq2/microserv/internal"
)

type Service interface {
	Get(context.Context, internal.Filter) (internal.Transfer, error)
	GetMany(context.Context, ...internal.Filter) (<-chan internal.Transfer, error)
	Status(context.Context, uint64) (internal.Status, error)
	Transfer(context.Context, *internal.Transfer) (string, error)
	ServiceStatus(ctx context.Context) (int, error)
}
