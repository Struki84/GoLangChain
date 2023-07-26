package memory

import "github.com/tmc/langchaingo/schema"

type PersistentChatMessageHistory struct {
	messages []schema.ChatMessage
	store    schema.ChatMessageHistoryStore
}

var _ schema.ChatMessageHistory = &PersistentChatMessageHistory{}

func NewPersistentChatMessageHistory(options ...PersistentChatMessageHistoryOption) *PersistentChatMessageHistory {
	return applyPersistentChatOptions(options...)
}

func (h *PersistentChatMessageHistory) GetSessionID() string {
	return h.store.GetSessionID()
}

func (h *PersistentChatMessageHistory) SetSessionID(id string) {
	h.store.SetSessionID(id)
}

func (h *PersistentChatMessageHistory) Messages() ([]schema.ChatMessage, error) {
	msgs, err := h.store.GetMessages()
	if err != nil {
		return nil, err
	}

	h.messages = msgs
	return h.messages, nil
}

func (h *PersistentChatMessageHistory) AddAIMessage(text string) error {
	msg := schema.AIChatMessage{Content: text}

	err := h.store.AddMessage(msg)
	if err != nil {
		return err
	}

	h.messages = append(h.messages, msg)
	return nil
}

func (h *PersistentChatMessageHistory) AddUserMessage(text string) error {
	msg := schema.HumanChatMessage{Content: text}

	err := h.store.AddMessage(msg)
	if err != nil {
		return err
	}

	h.messages = append(h.messages, msg)
	return nil
}

func (h *PersistentChatMessageHistory) AddMessage(message schema.ChatMessage) error {
	err := h.store.AddMessage(message)
	if err != nil {
		return err
	}

	h.messages = append(h.messages, message)
	return nil
}

func (h *PersistentChatMessageHistory) SetMessages(msgs []schema.ChatMessage) error {
	err := h.store.SetMessages(msgs)
	if err != nil {
		return err
	}
	h.messages = msgs
	return nil
}

func (h *PersistentChatMessageHistory) Clear() error {
	err := h.store.ClearMessages()
	if err != nil {
		return err
	}

	h.messages = make([]schema.ChatMessage, 0)
	return nil
}
