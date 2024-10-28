package domain

type BaseResponse struct {
	Message    string `json:"message"`
	ApiVersion string `json:"api_version"`
}

type BaseResponseWithData[C any] struct {
	BaseResponse
	Data C `json:"data"`
}

type BaseResponseWithPaginatedData[C any] struct {
	BaseResponse
	Total      int64 `json:"total"`
	PageNumber int   `json:"page_number"`
	PerPage    int   `json:"per_page"`
	Data       C     `json:"data"`
}
