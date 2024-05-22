//go:build go1.20
// +build go1.20

package context2

import (
	"context"
	"time"
)

func WithCancelCauseDetach(parent context.Context) (ctx context.Context, cancel context.CancelCauseFunc) {
	ctx2 := Detach(parent)
	return context.WithCancelCause(ctx2)
}

func WithDeadlineCauseDetach(parent context.Context, d time.Time, cause error) (context.Context, context.CancelFunc) {
	ctx2 := Detach(parent)
	return context.WithDeadlineCause(ctx2, d, cause)
}

func WithTimeoutCauseDetach(parent context.Context, timeout time.Duration, cause error) (context.Context, context.CancelFunc) {
	ctx2 := Detach(parent)
	return context.WithTimeoutCause(ctx2, timeout, cause)
}
