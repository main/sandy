# sandy
Golang library for notifications about slow code execution

# usage
	maxSilenceTime := 5 * time.Second
	maxOperationTime := 2 * time.Second
	instanceName := "Auto PG worker"

	sandy := app.New(context.Background(), app.Options{
		MaxSilenceTime:   maxSilenceTime,
		MaxOperationTime: maxOperationTime,
		MailerOptions: mailer.Options{
			InstanceName: "Sandy main.go",
			TemplateEmailSilenceMaxTimeExceeded: fmt.Sprintf(`
                Instance %s has reached maximum silence timeout,
				which is %s. Number of photos in handle {{.NumberOfPhotos}}`,instanceName, maxSilenceTime.String()),
			TemplateEmailOperationMaxTimeExceeded: fmt.Sprintf(`
                Instance %s has reached maximum operation timeout,
				which is %s. Number of photos in handle {{.NumberOfPhotos}}`, instanceName, maxOperationTime.String()),
			SenderEmail: "example@example.com",
			SenderName:  "Rachel",
			Receivers:   []string{"vannov@gmail.com"},
			SendGridKey: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		},
	})

	args := make(map[string]string)
	args["NumberOfPhotos"] = "213"

	sandy.OperationStarted(args)

    // ...
	// do some job
    // ...

	sandy.OperationFinished()


