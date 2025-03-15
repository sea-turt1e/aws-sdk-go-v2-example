# aws-sdk-go-v2-example
Example Code of [aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2)

## Copy .env.example to .env and fill in the values
```shell
cp .env.example .env
```

## Example Code
- [getPromptFromPromptManagement.go](bedrock/getPromptFromPromptManagement/main.go)
  - Get prompt from Amazon Bedrock Prompt Management Service

- [postClaudeWithToolUse.go](bedrock/postClaudeWithToolUse/main.go)
  - Get LLM Output of Claude with ToolUse From Amazon Bedrock.  
  - ToolUse is also called Function Calling.

