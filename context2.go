// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package xcontext is a package to offer the extra functionality we need
// from contexts that is not available from the standard context package.

// guonaihong 新增WithXXX相关函数
package context2

import (
	"context"
	"time"
)

// Detach returns a context that keeps all the values of its parent context
// but detaches from the cancellation and error handling.
func Detach(ctx context.Context) context.Context { return detachedContext{ctx} }

type detachedContext struct{ parent context.Context }

func (v detachedContext) Deadline() (time.Time, bool)       { return time.Time{}, false }
func (v detachedContext) Done() <-chan struct{}             { return nil }
func (v detachedContext) Err() error                        { return nil }
func (v detachedContext) Value(key interface{}) interface{} { return v.parent.Value(key) }

func WithCancelDetach(parent context.Context) (ctx context.Context, cancel context.CancelFunc) {
	ctx2 := Detach(parent)
	return context.WithCancel(ctx2)
}

func WithDeadlineDetach(parent context.Context, d time.Time) (context.Context, context.CancelFunc) {
	ctx2 := Detach(parent)
	return context.WithDeadline(ctx2, d)

}

func WithTimeoutDetach(parent context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	ctx2 := Detach(parent)
	return context.WithTimeout(ctx2, timeout)
}
