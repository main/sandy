package watch_dog

import (
	"context"
	"time"
)

func Watch(ctx context.Context, limit time.Duration, doneAfterTimeout, onTimeout func()) context.CancelFunc {
	cont, cancel := context.WithCancel(ctx)
	startTime := time.Now()
	go func() {
		ch := time.After(limit)

		select {
		case <-cont.Done():
			if time.Now().Add(-1 * limit).After(startTime) {
				doneAfterTimeout()
			}
		case <-ch:
			onTimeout()
		}
	}()

	return cancel
}
