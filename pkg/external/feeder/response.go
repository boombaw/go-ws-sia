package feeder

type FeederResponse struct {
	ErrorCode int                    `json:"error_code"`
	ErrorDesc string                 `json:"error_desc"`
	Data      map[string]interface{} `json:"data"`
}

// FeederResponses is a struct for feeder response for data array
type FeederResponses struct {
	ErrorCode int                      `json:"error_code"`
	ErrorDesc string                   `json:"error_desc"`
	Data      []map[string]interface{} `json:"data"`
}
