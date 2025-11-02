package response

type APIResponse struct {
    Status  bool        `json:"status"`
    Code    int         `json:"code"`
    Message string      `json:"message"`
    Data    interface{} `json:"data"`
}

func Success(code int, message string, data interface{}) APIResponse {
    return APIResponse{
        Status:  true,
        Code:    code,
        Message: message,
        Data:    data,
    }
}

func Error(code int, message string, data interface{}) APIResponse {
    return APIResponse{
        Status:  false,
        Code:    code,
        Message: message,
        Data:    data,
    }
}
