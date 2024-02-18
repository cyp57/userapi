package model

type Response struct {
	Result     interface{} `json:"result"`
	Status     bool        `json:"status"`
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
}
