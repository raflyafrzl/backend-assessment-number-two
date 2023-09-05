package model

type ResponseFailWeb struct {
	Error      any    `json:"error"`
	StatusCode int    `json:"status_code"`
	Status     string `json:"status"`
}
