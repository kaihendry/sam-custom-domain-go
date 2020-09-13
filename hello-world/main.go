package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/apex/log"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// Doesn't work
	// cfg, err := external.LoadDefaultAWSConfig(external.WithSharedConfigProfile("mine"))
	// Works
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		log.WithError(err).Fatal("setting up credentials")
	}
	cfg.Region = "ap-southeast-1"

	client := sns.New(cfg)
	req := client.PublishRequest(&sns.PublishInput{
		TopicArn: aws.String(os.Getenv("SNS_TOPIC_ARN")),
		Message:  aws.String(fmt.Sprintf("time: %s", time.Now())),
	})

	_, err = req.Send(context.Background())
	if err != nil {
		log.WithError(err).Fatal("unable to send SNS")
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "text/plain",
		},
		Body: "Hello World2",
	}, nil
}

func main() {
	lambda.Start(handler)
}
