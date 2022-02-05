package utils

type ErrorMess struct {
	Severity string `json:"severity"`
	Message  string `json:"message"`
}

type ApiError struct {
	Message string `json:"message"`
}
type ApiMessage struct {
	Message  string `json:"message"`
	Response string `json:"response"`
}