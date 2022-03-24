package helper

type meta struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type response struct {
	Meta meta        `json:"meta"`
	Data interface{} `json:"data"`
}

func ResponseAPI(message string, status int, data interface{}) response {
	meta := meta{
		Message: message,
		Status:  status,
	}
	response := response{
		Meta: meta,
		Data: data,
	}
	return response
}
