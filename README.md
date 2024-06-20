# go-hugging-api

This is a simple API that connect to huggingface.co API to get the model and classify the text.

## How to run
- Clone the repository
- Set the environment variable `HF_API_TOKEN` with your huggingface.co API key
- Run the command `go run main.go`

## Dependencies
- Redis
- Ratelimiter
- Gin