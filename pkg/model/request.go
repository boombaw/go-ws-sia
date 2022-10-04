package model

// MessageObject Basic chat message object
type MessageObject struct {
	Data  string `json:"data"`
	From  string `json:"from"`
	Event string `json:"event"`
	To    string `json:"to"`
	SID   string `json:"uuid"`
}

type FeederParams struct {
	Token string                 `json:"token"`
	Sms   string                 `json:"sms"`
	Data  map[string]interface{} `json:"data"`
}
