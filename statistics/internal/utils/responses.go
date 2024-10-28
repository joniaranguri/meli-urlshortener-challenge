package utils

import (
	"github.com/joniaranguri/meli-urlshortener-challenge/statistics/internal/core/domain"
	"github.com/joniaranguri/meli-urlshortener-challenge/statistics/internal/utils/constants"
)

func GetResponse(message string) domain.BaseResponse {
	return domain.BaseResponse{
		Message:    message,
		ApiVersion: constants.ApiVersion,
	}
}

func GetResponseWithData(message string, data any) domain.BaseResponseWithData[any] {
	return domain.BaseResponseWithData[any]{
		BaseResponse: GetResponse(message),
		Data:         data,
	}
}
