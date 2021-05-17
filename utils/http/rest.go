package http

// RestResponse : REST response
type RestResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}
