package usecase

import (
	"context"
)

type redisCartlist struct {
	MapPID map[int64]int64
	MapSID map[int64]int64
}

func (req *RequestFrontEndCartlist) validateInput(ctx context.Context) []string {
	return []string{}
}

func (resp *ResponseCartlist) setInitialResponse(isFromExternalEndpoint bool) {
	resp.ListShopID = nil
}

func (resp *ResponseCartlist) setResponseInputValidationError() {
	// set response value
}

func (resp *ResponseCartlist) setResponseTooManyRequest() {
	// set response value
}

type RequestFrontEndCartlist struct {
	productID int
	qty       int
	shopID    int
}

type ResponseCartlist struct {
	ListProductID []int
	ListShopID    []int
}
