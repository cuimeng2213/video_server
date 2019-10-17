package defs

// error handle struct
type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrorResponse struct {
	HttpSC int
	Error  Err
}

var (
	ErrorRequestBodyParseFailed = ErrorResponse{
		HttpSC: 400,
		Error: Err{
			Error:     "Request body is not correct",
			ErrorCode: "001",
		},
	}
	ErrorNotAUthUser = ErrorResponse{
		HttpSC: 401,
		Error: Err{
			Error:     "User Authendtication Failed",
			ErrorCode: "002",
		},
	}
)
