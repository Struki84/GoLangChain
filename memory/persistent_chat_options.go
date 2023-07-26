package memory

import "github.com/tmc/langchaingo/schema"

// PersistentChatMessageHistoryOption is a function for creating new persistent chat message history
// with other then the default values.
type PersistentChatMessageHistoryOption func(m *PersistentChatMessageHistory)

// WithMessages is an option for NewPersistentChatMessageHistory for adding
// previous messages to the history.
func WithMessages(previousMessages []schema.ChatMessage) PersistentChatMessageHistoryOption {
	return func(m *PersistentChatMessageHistory) {
		m.messages = append(m.messages, previousMessages...)
	}
}

// WithDBStore is an option for NewPersistentChatMessageHistory for adding
// db store manager.
func WithDBStore(dbstore schema.ChatMessageHistoryStore) PersistentChatMessageHistoryOption {
	return func(m *PersistentChatMessageHistory) {
		m.store = dbstore
	}
}

func applyPersistentChatOptions(options ...PersistentChatMessageHistoryOption) *PersistentChatMessageHistory {
	h := &PersistentChatMessageHistory{
		messages: make([]schema.ChatMessage, 0),
	}

	for _, option := range options {
		option(h)
	}

	return h
}
