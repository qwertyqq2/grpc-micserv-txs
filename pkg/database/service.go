package database

import (
	"context"

	"github.com/qwertyqq2/microserv/internal"
)

type Service interface {
	AssetsDB
	UsersDB

	ServiceStatus(ctx context.Context) (int, error)
}

type UsersDB interface {
	GetUser(context.Context, *internal.UserHand) (*internal.User, error)
	RemoveUser(context.Context, ...*internal.UserHand) (string, error)
	SetBalance(context.Context, *internal.UserHand, uint64) (string, error)
	SetNickname(context.Context, *internal.UserHand, string) (string, error)
}

type AssetsDB interface {
	AddTx(context.Context, *internal.Transaction) (string, error)
	GetTx(context.Context, ...internal.Filter) (*internal.Transaction, error)
}
