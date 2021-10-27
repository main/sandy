package watch_dog

import (
	"context"
	"time"
)

func Watch(ctx context.Context, limit time.Duration, doneAfterTimeout func(), onTimeout func()) context.CancelFunc {
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		select {
		case <-ctx.Done():
		case <-time.After(limit):
			onTimeout()
			<-ctx.Done()
			doneAfterTimeout()
		}
	}()

	return cancel
}
