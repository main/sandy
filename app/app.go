package app

import (
	"context"
	"github.com/main/sandy/mailer"
	"github.com/main/sandy/watch_dog"
	"log"
	"time"
)

type App struct {
	ctx     context.Context
	options Options

	cancelMaxSilenceWatchDog   context.CancelFunc
	cancelMaxOperationWatchDog context.CancelFunc

	mailer *mailer.Mailer
}

type Options struct {
	MaxSilenceTime   time.Duration
	MaxOperationTime time.Duration

	MailerOptions mailer.Options
}

func New(ctx context.Context, opts Options) *App {
	return &App{
		ctx:     ctx,
		options: opts,

		mailer: mailer.NewMailer(opts.MailerOptions),
	}
}

func (a *App) OperationStarted(templateArgs map[string]string) {
	a.cancelMaxSilenceWatchDog = watch_dog.Watch(a.ctx, a.options.MaxSilenceTime, func() {}, func() {
		log.Println("Max silence func")
		if err := a.mailer.SendMaxSilenceEmails(templateArgs); err != nil {
			//TODO: use logger
			log.Println("SendMaxSilenceEmails err ", err)
		}
	})

	a.cancelMaxOperationWatchDog = watch_dog.Watch(a.ctx, a.options.MaxOperationTime, func() {
		log.Println("Max operation func")
		if err := a.mailer.SendMaxOperationEmails(templateArgs); err != nil {
			//TODO: use logger
			log.Println("SendMaxOperationEmails err ", err)
		}
	}, func() {})
}

func (a *App) OperationFinished() {
	a.cancelMaxSilenceWatchDog()
	a.cancelMaxOperationWatchDog()
}
