package app

import (
	"context"
	"github.com/main/sandy/watch_dog"
	"log"
	"time"
)

type App struct {
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
	app := &App{}

	return app, nil
}

func (a *App) OperationStarted(args ...string) {
	a.cancelMaxSilenceWatchDog = watch_dog.Watch(context.TODO(), a.options.MaxSilenceTime, func() {
		log.Println("func1")
	}, func() {
		log.Println("func2")
	})

	a.cancelMaxOperationWatchDog = watch_dog.Watch(context.TODO(), a.options.MaxOperationTime, func() {
		log.Println("func1")
	}, func() {
		log.Println("func2")
	})
}

func (a *App) OperationFinished() {
	a.cancelMaxSilenceWatchDog()
	a.cancelMaxOperationWatchDog()
}
