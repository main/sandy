# sandy
Golang library for notifications about slow code execution

# usage
    go get 	sandy github.com/main/sandy
    
    import sandy "github.com/main/sandy/app"

	maxSilenceTime := 5 * time.Second
	maxOperationTime := 2 * time.Second
	instanceName := "Auto PG worker"

	snd := sandy.New(context.Background(), app.Options{
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

	snd.OperationStarted(args)

    // ...
	// do some job
    // ...

	snd.OperationFinished()


