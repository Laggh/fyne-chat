package main

type Message struct {
	MessageID   int    `json:"message_id"`
	SenderID    string `json:"sender_id"`
	TextContent string `json:"text_content,omitempty"` // omitempty: ignora se estiver vazio
	ImageURL    string `json:"image_url,omitempty"`
	Timestamp   int64  `json:"timestamp"` // Use int64 para Unix Epoch (milissegundos/segundos)
}
