package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/disintegration/imaging"
	"github.com/joho/godotenv"
)

type ReqBody struct {
	Bucket    string `json:"bucket"`
	ObjectKey string `json:"objectKey"`
	SaveName  string `json:"saveName"`
}

func main() {
	// get arguments from command line
	args := os.Args
	if len(args) > 1 && args[1] == "local" {
		// set env from .env file
		err := godotenv.Load()
		if err != nil {
			fmt.Println("Error in loading .env file")
		}
		// Run the server locally
		RunLocally(Handler)
	} else {
		// Run the Lambda function
		lambda.Start(Handler)
	}
}

// handler is the Lambda function handler which receives body of the request
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var reqBody ReqBody

	// Get the bucket name and object key from the request body
	err := json.Unmarshal([]byte(request.Body), &reqBody)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: "Error in unmarshalling request body", StatusCode: 500}, nil
	}

	bucket := reqBody.Bucket
	objectKey := reqBody.ObjectKey
	saveName := reqBody.SaveName

	// get s3 client
	client := getS3()

	// Get the object from the S3 bucket to compress
	file, err := getObject(client, bucket, objectKey)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: "Error in getting object from S3", StatusCode: 500}, nil
	}

	// compress the image
	compressedFile, err := compressImage(file)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: "Error in compressing image", StatusCode: 500}, nil
	}

	// Save the compressed image to the S3 bucket
	err = saveObject(client, bucket, saveName, compressedFile)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: "Error in saving object to S3", StatusCode: 500}, nil
	}

	return events.APIGatewayProxyResponse{Body: "Image compressed successfully", StatusCode: 200}, nil
}

// compressImage compresses the image and returns the compressed image
func compressImage(file []byte) ([]byte, error) {
	// Decode the image
	img, err := imaging.Decode(bytes.NewReader(file))
	if err != nil {
		fmt.Println("Error in decoding image")
		return nil, fmt.Errorf("error in decoding image")
	}

	// Resize the image to width 128px
	img = imaging.Resize(img, 128, 128, imaging.Lanczos)

	// Encode the image to JPEG
	buf := new(bytes.Buffer)
	err = imaging.Encode(buf, img, imaging.JPEG)
	if err != nil {
		fmt.Println("Error in encoding image")
		return nil, fmt.Errorf("error in encoding image")
	}

	return buf.Bytes(), nil
}

// getObject gets the object from the S3 bucket and returns a byte array
func getObject(client *s3.Client, bucket string, objectKey string) ([]byte, error) {
	// Create a new GetObjectInput
	input := &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &objectKey,
	}

	// Get the object from the S3 bucket
	result, err := client.GetObject(context.Background(), input)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("error in getting object from S3")
	}

	file, _ := io.ReadAll(result.Body)

	return file, nil
}

// saveObject saves the object to the S3 bucket
func saveObject(client *s3.Client, bucket string, objectKey string, file []byte) error {
	// Create a new PutObjectInput
	input := &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &objectKey,
		Body:   bytes.NewReader(file),
	}

	// Put the object to the S3 bucket
	_, err := client.PutObject(context.Background(), input)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("error in putting object to S3")
	}

	return nil
}

// awsHandler creates a new S3 client
func getS3() (awsS3Client *s3.Client) {

	creds := credentials.NewStaticCredentialsProvider(os.Getenv("ACCESS_KEY_AWS"), os.Getenv("SECRET_KEY_AWS"), "")

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(creds), config.WithRegion(os.Getenv("REGION_AWS")))
	if err != nil {
		fmt.Println("Error in loading AWS Config")
		return nil
	}

	// Create an S3 service client
	client := s3.NewFromConfig(cfg)

	return client
}
