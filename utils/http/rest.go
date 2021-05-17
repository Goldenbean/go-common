package http

// RestResponse : REST response
type RestResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

func SuccessRestResponse() *RestResponse {
	return &RestResponse{
		Code:    200,
		Message: "sucess",
		Success: true,
	}
}

func FailedRestResponse() *RestResponse {
	return &RestResponse{
		Code:    200,
		Message: "failed",
		Success: false,
	}
}
