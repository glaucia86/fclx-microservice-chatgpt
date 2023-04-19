/**
 * file: internal/domain/entity/model.go
 * date: 04/18/2023
 * description: file responsible for the model entity
 * author: Glaucia Lemos <Twitter: @glaucia_lemos86>
 */

package entity

type Model struct {
	Name      string
	MaxTokens int
}

func NewModel(name string, maxTokens int) *Model {
	return &Model{
		Name:      name,
		MaxTokens: maxTokens,
	}
}

func (m *Model) GetModelName() string {
	return m.Name
}

func (m *Model) GetModelMaxTokens() int {
	return m.MaxTokens
}
