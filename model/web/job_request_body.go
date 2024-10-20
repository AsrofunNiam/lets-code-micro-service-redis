package web

type JobsQueueCreateRequest struct {
	// Fields
	Method  string `json:"method" validate:"required,min=1"`
	URL     string `json:"url" validate:"required,min=1"`
	Payload string `json:"payload" validate:"required,min=1"`
	Key     string `json:"key" validate:"required,min=1"`
}
