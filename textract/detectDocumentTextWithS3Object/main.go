// Description: Detect Text of S3 Object using AWS Textract

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/textract"
	"github.com/aws/aws-sdk-go-v2/service/textract/types"
	"github.com/joho/godotenv"
)

func init() {
	// Load Environment Variables
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func detectDocumentTextWithS3Object() {
	// Load the SDK's configuration from the environment
	sdkConfig, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Create a Texttract client
	textractClient := textract.NewFromConfig(sdkConfig)

	// Get Document Text from S3 Object
	detectDocumentTextDetectionOutput, err := textractClient.DetectDocumentText(context.Background(), &textract.DetectDocumentTextInput{
		Document: &types.Document{
			S3Object: &types.S3Object{
				Bucket: aws.String(os.Getenv("BUCKET_NAME")),
				Name:   aws.String(os.Getenv("OBJECT_NAME")),
			},
		},
	})
	if err != nil {
		log.Fatalf("unable to detect document text, %v", err)
	}
	detectedText := getTextFromTextractOutput(detectDocumentTextDetectionOutput)
	fmt.Println(detectedText)
}

func getTextFromTextractOutput(output *textract.DetectDocumentTextOutput) string {
	text := ""
	for _, block := range output.Blocks {
		if block.BlockType == types.BlockTypeLine {
			text += *block.Text + "\n"
		}
	}
	return text
}

func main() {
	detectDocumentTextWithS3Object()
}
