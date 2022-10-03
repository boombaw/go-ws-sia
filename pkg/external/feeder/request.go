package feeder

type GetPayload struct {
	Act    string `json:"act"`
	Token  string `json:"token"`
	Filter string `json:"filter"`
	Limit  string `json:"limit"`
	Order  string `json:"order"`
}

type Token struct {
	Act      string `json:"act"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type PutPayload struct {
	Act    string                 `json:"act"`
	Token  string                 `json:"token"`
	Key    map[string]interface{} `json:"key"`
	Record map[string]interface{} `json:"record"`
}

type Key struct {
	Data map[string]interface{} `json:"key"`
}

type Record struct {
	Data map[string]interface{} `json:"record"`
}
