package domain

type HTTPResponse[T any] struct {
	Body     *T       `json:"body"`
	Severity Severity `json:"severity"`
	Message  string   `json:"message"`
}
