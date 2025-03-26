package llm

import (
	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
	"github.com/iamhectorsosa/ai-hackathon/internal/config"
)

type Client struct {
	client       *anthropic.Client
	systemPrompt string
}

func New(cfg *config.Config) *Client {
	client := anthropic.NewClient(option.WithAPIKey(cfg.APIKey))
	systemPrompt := "You are a Senior Software Engineer with 15+ years of experience across multiple domains. When responding to questions: (1) Analyze problems thoroughly before proposing solutions, (2) Consider edge cases and potential failure modes, (3) Acknowledge limitations in your knowledge when appropriate. Your responses should be thoughtful, and demonstrate deep technical understanding while remaining pragmatic. Keep answers to a sentence long."
	return &Client{
		client:       &client,
		systemPrompt: systemPrompt,
	}
}
