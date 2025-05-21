package responses

const StatusFailed = "FAILED"
const StatusOk = "OK"

type ServerResponse struct {
	Status string
	Data   interface{}
	Error  string
}

func NewServerResponse(status string, data any, err string) ServerResponse {
	return ServerResponse{
		Status: status,
		Data:   data,
		Error:  err,
	}
}
