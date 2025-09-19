package rest_err

type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int64    `json:"code"`
	Causes  []Causes `json:"causes,omitempty"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewRestError(message string, err string, code int64, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func (*RestErr) Error() string {
	return r.Message
}

func NewBadRequestError(message string) *RestErr {
	return NewRestError(message, "bad_request", 400, nil)
}
func NewBadRequestValidationError(message string, causes []Causes) *RestErr {
	return NewRestError(message, "bad_request", 400, causes)
}
func NewIntenalServerError(message string) *RestErr {
	return NewRestError(message, "internal_server_error", 500, nil)
}
func NewNotFoundError(message string) *RestErr {
	return NewRestError(message, "not_found", 404, nil)
}
