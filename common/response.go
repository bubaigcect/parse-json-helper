package response

type Response struct {
	Status       bool        `json:"status"`
	Code         int         `json:"code"`
	ResponseCode int         `json:"response_code,omitempty"`
	Message      string      `json:"message,omitempty"`
	Data         interface{} `json:"data,omitempty"`
}
