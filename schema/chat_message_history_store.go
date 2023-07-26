package schema

type ChatMessageHistoryStore interface {
	// SetSessionId, method for setting user session id
	SetSessionID(id string)

	// GetSessionId, method for getting user session id
	GetSessionID() string

	// AddMessage, method for storing a message in the DB store
	AddMessage(msg ChatMessage) error

	// SetMessages, method for replacing existing messages in the DB store
	SetMessages(msgs []ChatMessage) error

	// GetMessages, Convinience method for getting messages from db store
	GetMessages() ([]ChatMessage, error)

	// ClearMessages, method for clearing messages in the DB store
	ClearMessages() error
}
