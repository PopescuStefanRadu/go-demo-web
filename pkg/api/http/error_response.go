package http

type ErrorResponse struct {
	GlobalErrors []string            `json:"global_errors,omitempty"`
	FieldErrors  map[string][]string `json:"field_errors,omitempty"`
}
