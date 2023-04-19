/**
 * file: internal/domain/entity/chat.go
 * date: 04/18/2023
 * description: file responsible for the creation of the entity of the application
 * author: Glaucia Lemos <Twitter: @glaucia_lemos86>
 */

package entity

// Reference: https://platform.openai.com/docs/api-reference/chat/create

type ChatConfig struct {
	Model            *Model
	Temperature      float32
	TopP             float32
	N                int
	Stop             []string
	MaxTokens        int
	PresencePenalty  float32
	FrequencyPenalty float32
}

type Chat struct {
	ID                   string
	UserID               string
	InitialSystemMessage *Message
	Messages             []*Message
	ErasedMessages       []*Message
	Status               string
	TokenUsage           int
	Config               *ChatConfig
}
