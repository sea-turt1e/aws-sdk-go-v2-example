# aws-sdk-go-v2-example
Example Code of [aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2)

## Copy .env.example to .env and fill in the values
```shell
cp .env.example .env
```

## Example Code
- [getPromptFromPromptManagement](bedrock/getPromptFromPromptManagement/main.go)
  - Get prompt from Amazon Bedrock Prompt Management Service

- [postClaudeWithToolUse](bedrock/postClaudeWithToolUse/main.go)
  - Get LLM Output of Claude with ToolUse From Amazon Bedrock.  
  - ToolUse is also called Function Calling.

- [detectDocumentTextWithS3Object](textract/detectDocumentTextWithS3Object/main.go)
  - Detect Document Text with S3 Object using Amazon Textract

- [analyzeDocumentTextWithS3Object](textract/analyzeDocumentTextWithS3Object/main.go)
  - Analyze Document Text with S3 Object using Amazon Textract