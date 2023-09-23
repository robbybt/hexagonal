package usecase

import (
	"context"
)

type redisAtc struct {
	MapPID map[int64]int64
	MapSID map[int64]int64
}

func (req *RequestFrontEndATC) validateInput(ctx context.Context) []string {
	return []string{}
}

func (resp *ResponseATC) setInitialResponse(isFromExternalEndpoint bool) {
	resp.ListShopID = nil
}

func (resp *ResponseATC) setResponseInputValidationError() {
	// set response value
}

func (resp *ResponseATC) setResponseTooManyRequest() {
	// set response value
}

type RequestFrontEndATC struct {
	productID int
	qty       int
	shopID    int
}

type ResponseATC struct {
	ListProductID []int
	ListShopID    []int
}
