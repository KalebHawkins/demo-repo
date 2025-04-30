package main

import "github.com/aws/aws-cdk-go/awscdk/v2"

func main() {
	app := awscdk.NewApp(nil)

	NewLambdaStack(app, "demo-projectabc-lambda-stack", nil)

	app.Synth(nil)
}

