package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagent"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagent/types"
	"github.com/joho/godotenv"
)

func init() {
	// Load Environment Variables
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// Get Prompt From Bedrock Prompt Management
	// This is a simple example of how to get a prompt from the Bedrock Prompt Management API.
	// This example assumes that you have a prompt with the ID of "promptId" in your Bedrock Prompt Management.

	// Load the model ID and prompt ID from the environment
	promptId := os.Getenv("PROMPT_ID")

	// Load the SDK's configuration from the environment
	sdkConfig, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Create a new Prompt Management client
	bedrockAgent := bedrockagent.NewFromConfig(sdkConfig)
	getPromptOutput, err := bedrockAgent.GetPrompt(context.Background(), &bedrockagent.GetPromptInput{
		PromptIdentifier: &promptId,
	},
	)
	if err != nil {
		log.Fatalf("unable to get prompt, %v", err)
	}

	promptVariant := getPromptOutput.Variants[0]
	promptTemplateType := promptVariant.TemplateType
	if promptTemplateType == "CHAT" {
		prompt := promptVariant.TemplateConfiguration.(*types.PromptTemplateConfigurationMemberChat)
		fmt.Println("Prompt: ", prompt.Value.System[0].(*types.SystemContentBlockMemberText).Value)
		for i, message := range prompt.Value.Messages {
			fmt.Printf("Message%s: %s \n", strconv.Itoa(i), message.Content[0].(*types.ContentBlockMemberText).Value)
		}
	} else if promptTemplateType == "TEXT" {
		prompt := promptVariant.TemplateConfiguration.(*types.PromptTemplateConfigurationMemberText)
		fmt.Printf("Prompt: %s", *prompt.Value.Text)
	} else {
		log.Fatalf("unable to get prompt, %v", err)
	}
}
