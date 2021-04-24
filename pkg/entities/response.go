package entities

// Response is a simple response from an HTTP service.
// Body is assumed to be an object that is json serializable.
type Response struct {
	Code     int         `json:"code,omitempty"`
	Body     interface{} `json:"body,omitempty"`
	Title    string      `json:"title,omitempty"`
	Message  string      `json:"message,omitempty"`
	Instance string      `json:"instance,omitempty"`
}

// Problem defines the Problem JSON type defined by RFC 7807 - media type
// application/problem+json.
// It should be the expected error response for all APIs.
type Problem struct {
	Type     string `json:"type,omitempty"`
	Title    string `json:"title,omitempty"`
	Status   int    `json:"status,omitempty"`
	Detail   string `json:"detail,omitempty"`
	Instance string `json:"instance,omitempty"`
}
