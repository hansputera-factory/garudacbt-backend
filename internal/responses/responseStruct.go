package responses

type baseResponse struct {
	Message          string `json:"message"`
	Data             any    `json:"data"`
	Ok               bool   `json:"ok"`
	ValidationErrors any    `json:"validation_errors"`
}
