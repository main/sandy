package app

import (
	"context"
	"github.com/main/sandy/watch_dog"
	"time"
)

type App struct {
	ctx     context.Context
	options Options

	cancelMaxSilenceWatchDog   context.CancelFunc
	cancelMaxOperationWatchDog context.CancelFunc
}

type Options struct {
	MaxSilenceTime   time.Duration
	MaxOperationTime time.Duration

	TextSilenceMaxTimeExceeded   func(...string) (string, error)
	TextOperationMaxTimeExceeded func(...string) (string, error)

	SenderEmail string
	SenderName  string
	Receivers   []string

	SendGridKey string
}

func New(ctx context.Context, opts Options) (*App, error) {
	app := &App{
		ctx:     ctx,
		options: opts,
	}

	return app, nil
}

func (a *App) OperationStarted(args ...string) {
	a.cancelMaxSilenceWatchDog = watch_dog.Watch(a.ctx, a.options.MaxSilenceTime, func() {}, func() {
		//TODO: send max silence email
	})

	a.cancelMaxOperationWatchDog = watch_dog.Watch(a.ctx, a.options.MaxOperationTime, func() {
		//TODO: send timeout email
	}, func() {})
}

func (a *App) OperationFinished() {
	a.cancelMaxSilenceWatchDog()
	a.cancelMaxOperationWatchDog()
}
