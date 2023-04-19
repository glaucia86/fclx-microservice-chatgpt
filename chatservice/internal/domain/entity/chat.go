/**
 * file: internal/domain/entity/chat.go
 * date: 04/18/2023
 * description: file responsible for the creation of the entity of the application
 * author: Glaucia Lemos <Twitter: @glaucia_lemos86>
 */

package entity

type ChatConfig struct {
	Model *Model
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
