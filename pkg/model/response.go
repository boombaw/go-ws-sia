package model

type MessageResponse struct {
	Event   string `json:"event"`
	Message string `json:"message"`
	ID      string `json:"uuid"`
}
