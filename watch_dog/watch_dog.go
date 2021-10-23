package watch_dog

import (
	"context"
	"time"
)

func Watch(ctx context.Context, limit time.Duration, done, onTimeout func()) context.CancelFunc {
	cont, cancel := context.WithCancel(ctx)

	go func() {
		ch := time.After(limit)

		select {
		case <-cont.Done():
			done()
		case <-ch:
			onTimeout()
		}
	}()

	return cancel
}
