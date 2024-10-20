package domain

type JobQueue struct {
	Method  string `json:"method"`
	URL     string `json:"url"`
	Payload string `json:"payload"`
}
