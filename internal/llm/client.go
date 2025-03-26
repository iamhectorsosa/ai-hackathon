package llm

import (
	"context"
	"strings"

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

func (c *Client) GetCompletion(ctx context.Context, prompt string) (string, error) {
	message, err := c.client.Messages.New(ctx, anthropic.MessageNewParams{
		Model:     anthropic.ModelClaude3_5HaikuLatest,
		MaxTokens: 10,
		System: []anthropic.TextBlockParam{
			{
				Type: "text",
				Text: c.systemPrompt,
			},
		},
		Messages: []anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock(prompt)),
		},
	})
	if err != nil {
		return "", err
	}

	var sb strings.Builder
	for _, content := range message.Content {
		if content.Type == "text" {
			sb.WriteString(content.Text)
		}
	}

	return sb.String(), nil
}
