package auth

import (
	"context"

	"github.com/qwertyqq2/microserv/internal"
)

type SessionID string

type Service interface {
	Auth(context.Context, string) (string, error)
	Rgstr(context.Context, *internal.UserHand) (string, error)

	SessionsDB
	ServiceStatus(ctx context.Context) (int, error)
}

type SessionsDB interface {
	NewSession(context.Context, *internal.UserHand) (SessionID, error)
	Get(context.Context, SessionID) (string, error)
	Remove(context.Context, SessionID) (string, error)
}

// func SessionsHook(ctx context.Context, sesID SessionID,
// 	serv Service, userHadn *internal.UserHand, proc func(ctx context.Context) (string, error)) (string, error) {
// 	if sid, err := serv.Get(ctx, sesID); err != nil || sid == "" {
// 		return "", ErrSessionNotFound
// 	}
// 	res, err := proc(ctx)
// 	if err != nil {
// 		return "", err
// 	}
// 	if _, err := serv.Remove(ctx, sesID); err != nil {
// 		return "", err
// 	}
// 	if _, err := serv.NewSession(ctx, userHadn); err != nil {
// 		return "", err
// 	}
// 	return res, nil
// }
