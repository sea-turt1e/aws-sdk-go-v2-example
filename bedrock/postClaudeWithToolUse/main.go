package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime/document"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime/types"
	"github.com/joho/godotenv"
)

func init() {
	// Load Environment Variables
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}
}

var toolName = "math_tool"

var itemProperties = map[string]interface{}{
	"formula": map[string]interface{}{
		"description": "Formula to be calculated",
		"type":        "string",
	},
	"answer": map[string]interface{}{
		"description": "the answer to the formula",
		"type":        "string",
	},
}

var toolProperties = map[string]interface{}{
	"item": map[string]interface{}{
		"type":       "array",
		"properties": itemProperties,
	},
}

var inpuSchema = map[string]interface{}{
	"type": "object",
	"properties": map[string]interface{}{
		"tool": toolProperties,
	},
}

func main() {
	// Get Prompt From Bedrock Prompt Management
	// This is a simple example of how to get a prompt from the Bedrock Prompt Management API.
	// This example assumes that you have a prompt with the ID of "promptId" in your Bedrock Prompt Management.

	// Load the model ID from the environment
	modelId := os.Getenv("MODEL_ID")

	// Load the SDK's configuration from the environment
	sdkConfig, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Create a new Bedrock Runtime client
	bedrockRuntime := bedrockruntime.NewFromConfig(sdkConfig)

	// Set up System Instruction
	systemInstruction := fmt.Sprintf("Use %s to get the sum of two numbers.", toolName)
	system := []types.SystemContentBlock{
		&types.SystemContentBlockMemberText{
			Value: systemInstruction,
		},
	}

	// Set up User Message
	input := "1 + 2"
	userMsg := types.Message{
		Role: types.ConversationRoleUser,
		Content: []types.ContentBlock{
			&types.ContentBlockMemberText{
				Value: input,
			},
		},
	}

	toolConfig := types.ToolConfiguration{
		Tools: []types.Tool{
			&types.ToolMemberToolSpec{
				Value: types.ToolSpecification{
					InputSchema: &types.ToolInputSchemaMemberJson{
						Value: document.NewLazyDocument(inpuSchema),
					},
					Name: &toolName,
				},
			},
		},
	}

	// Request to post a Claude with tool use
	output, err := bedrockRuntime.Converse(context.Background(), &bedrockruntime.ConverseInput{
		ModelId:    &modelId,
		Messages:   []types.Message{userMsg},
		System:     system,
		ToolConfig: &toolConfig,
	})
	if err != nil {
		log.Fatalf("unable to post Claude with tool use, %v", err)
	}

	response, _ := output.Output.(*types.ConverseOutputMemberMessage)
	responseContentBlock := response.Value.Content[0]
	text, _ := responseContentBlock.(*types.ContentBlockMemberText)
	fmt.Printf("Response: %s\n", text.Value)

	contentBlock := response.Value.Content[1]
	toolUseOutput, _ := contentBlock.(*types.ContentBlockMemberToolUse)
	toolUseOutputJson, err := toolUseOutput.Value.Input.MarshalSmithyDocument()
	if err != nil {
		log.Fatalf("unable to marshal tool use output, %v", err)
	}
	fmt.Printf("Tool Use Output: %s\n", toolUseOutputJson)
}
