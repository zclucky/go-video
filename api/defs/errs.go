package defs

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrorResponse struct {
	HttpSC int
	Error  Err
}
// 定义错误
var (
	ErrorRequestBodyParseFailed = ErrorResponse{
		HttpSC: 400,
		Error: Err{
			Error:     "Request body is not correct",
			ErrorCode: "0001",
		},
	}
	ErrorNotAuthUser = ErrorResponse{
		HttpSC: 401,
		Error: Err{
			Error:     "User authentication failed",
			ErrorCode: "0002",
		},
	}
)
