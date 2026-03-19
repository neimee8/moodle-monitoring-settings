package types

type ApiResponseData struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ApiResponse struct {
	Status       int
	ResponseData *ApiResponseData
}

func NewApiResponse() *ApiResponse {
	return &ApiResponse{
		Status: 200,
		ResponseData: &ApiResponseData{
			Data: nil,
		},
	}
}
