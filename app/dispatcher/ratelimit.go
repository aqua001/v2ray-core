// +build !confonly

package dispatcher

import (
	"context"

	"v2ray.com/core/common"
	"v2ray.com/core/common/buf"

	"golang.org/x/time/rate"
)

type RateLimitWriter struct {
	Writer  buf.Writer
	Limiter *rate.Limiter
	Ctx     context.Context
}

func (w *RateLimitWriter) WriteMultiBuffer(mb buf.MultiBuffer) error {
	w.Limiter.WaitN(w.Ctx, int(mb.Len()))
	return w.Writer.WriteMultiBuffer(mb)
}

func (w *RateLimitWriter) Close() error {
	return common.Close(w.Writer)
}

func (w *RateLimitWriter) Interrupt() {
	common.Interrupt(w.Writer)
}
