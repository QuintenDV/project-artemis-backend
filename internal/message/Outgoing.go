package message

type ErrorMessage struct {
	Message
	Content string `json:"message"`
}

type LobbyMessage struct {
	Message
	Lobby interface{} `json:"lobby"`
}

type ConnectedMessage struct {
	Message
}

type RegisteredMessage struct {
	Message
	Id int `json:"id"`
}

type HistoryMessage struct {
	Message
	Messages []*StateMessage `json:"messages"`
}

type CreatedMessage struct {
	Message
	Room int `json:"room"`
}

func NewErrorMessage(content string) *ErrorMessage {
	message := Message {
		Type: "error",
	}
	return &ErrorMessage {
		Message: message,
		Content: content,
	}
}

func NewLobbyMessage(lobby interface{}) *LobbyMessage {
	message := Message {
		Type: "lobby",
	}
	return &LobbyMessage {
		Message: message,
		Lobby: lobby,
	}
}

func NewConnectedMessage() *ConnectedMessage {
	message := Message {
		Type: "connected",
	}
	return &ConnectedMessage {
		Message: message,
	}
}

func NewRegisteredMessage(id int) *RegisteredMessage {
	message := Message {
		Type: "registered",
	}
	return &RegisteredMessage {
		Message: message,
		Id: id,
	}
}

func NewHistoryMessage(messages []*StateMessage) *HistoryMessage {
	message := Message {
		Type: "history",
	}
	return &HistoryMessage {
		Message: message,
		Messages: messages,
	}
}

func NewCreatedMessage(room int) *CreatedMessage {
	message := Message {
		Type: "created",
	}
	return &CreatedMessage {
		Message: message,
		Room: room,
	}
}
