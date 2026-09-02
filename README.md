# sandy

A small Go watchdog library that emails you when your code runs slower than expected.

Wrap a long-running operation with `OperationStarted` / `OperationFinished`, and sandy will send a notification (via SendGrid) if:

- **MaxOperationTime** is exceeded — the operation is taking too long overall;
- **MaxSilenceTime** is exceeded — the operation hasn't reported progress for too long.

Useful for background workers and batch jobs where a hang would otherwise go unnoticed.

## Install

```
go get github.com/main/sandy
```

## Usage

```go
package main

import (
	"context"
	"time"

	sandy "github.com/main/sandy/app"
	"github.com/main/sandy/mailer"
)

func main() {
	snd := sandy.New(context.Background(), sandy.Options{
		MaxSilenceTime:   5 * time.Second,
		MaxOperationTime: 2 * time.Second,
		MailerOptions: mailer.Options{
			InstanceName: "photo-worker",
			TemplateEmailSilenceMaxTimeExceeded: `
				Instance has reached the maximum silence timeout.
				Photos in progress: {{.NumberOfPhotos}}`,
			TemplateEmailOperationMaxTimeExceeded: `
				Instance has reached the maximum operation timeout.
				Photos in progress: {{.NumberOfPhotos}}`,
			SenderEmail: "alerts@example.com",
			SenderName:  "Sandy",
			Receivers:   []string{"you@example.com"},
			SendGridKey: "<your SendGrid API key>",
		},
	})

	// template arguments are rendered into the email body
	args := map[string]interface{}{"NumberOfPhotos": "213"}

	snd.OperationStarted(args)

	// ... do the work ...

	snd.OperationFinished()
}
```

Email templates use Go's `text/template` syntax; the arguments passed to `OperationStarted` are available inside the templates.

## How it works

`OperationStarted` launches two watchdog timers (see `watch_dog`). If either timer fires before `OperationFinished` cancels them, the corresponding email is sent through the `mailer` package. Cancelling the parent context stops the watchdogs as well.

## License

MIT — see [LICENSE](LICENSE).
