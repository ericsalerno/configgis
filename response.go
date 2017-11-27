package main

// Response envelope
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Key     string `json:"key"`
	Value   string `json:"value"`
}

// ValidResponse create a valid response
func ValidResponse(key string, value string) Response {
	r := Response{}
	r.Status = 200
	r.Message = "OK"
	r.Key = key
	r.Value = value

	return r
}

// ErrorResponse create a valid response
func ErrorResponse(err error) Response {
	r := Response{}
	r.Status = 500
	r.Key = ""
	r.Value = ""
	r.Message = err.Error()

	return r
}

//  FailResponse creates a failed usage response object
func FailResponse() Response {
	r := Response{}
	r.Status = 400
	r.Key = ""
	r.Value = ""
	r.Message = "Invalid operation. Use GET /get/<server>/<stage>/<key> or POST /set/<server>/<stage>/<key>"

	return r
}
