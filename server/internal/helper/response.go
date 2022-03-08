package helper

type Response struct {
	StatusCode  int         `json:"status"`
	Message 	string      `json:"message"`
	Data    	interface{} `json:"data"`
}

func BuildResponse(status int, message string, data interface{}) Response {
	res := Response{
		StatusCode: status,
		Message: 	message,
		Data:    	data,
	}
	return res
}

