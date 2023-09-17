package newprovider

import (
	"fmt"
	restriction2 "hexagonal/domain/entities/restriction"
)

type ValidateRestrictionResp struct {
	ValidateRestrictionResponse
}

type ValidateRestrictionResponse struct {
	Success      bool            `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message      string          `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	DataResponse []*DataResponse `protobuf:"bytes,3,rep,name=data_response,json=dataResponse,proto3" json:"data_response,omitempty"`
	MetaResponse []*MetaResponse `protobuf:"bytes,4,rep,name=meta_response,json=metaResponse,proto3" json:"meta_response,omitempty"`
}

type MetaResponse struct {
	RestrictionName string          `protobuf:"bytes,1,opt,name=restriction_name,json=restrictionName,proto3" json:"restriction_name,omitempty"`
	DataResponse    []*DataResponse `protobuf:"bytes,2,rep,name=data_response,json=dataResponse,proto3" json:"data_response,omitempty"`
	ErrorMessage    []string        `protobuf:"bytes,3,rep,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

type DataResponse struct {
	ProductId        int64  `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Status           string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	ShopId           int64  `protobuf:"varint,4,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
	PaymentGatewayId int64  `protobuf:"varint,5,opt,name=payment_gateway_id,json=paymentGatewayId,proto3" json:"payment_gateway_id,omitempty"`
	CouponId         string `protobuf:"bytes,8,opt,name=coupon_id,json=couponId,proto3" json:"coupon_id,omitempty"`
	UserId           int64  `protobuf:"varint,11,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	IsError          bool   `protobuf:"varint,12,opt,name=is_error,json=isError,proto3" json:"is_error,omitempty"`
}

func (repo *Repositories) ValidateRestrictionCategory(i restriction2.InputValidateRestrictionCategory) (restriction2.ValidateRestrictionResponse, error) {
	// restapicalls
	fmt.Println("rest API with", i)
	var resp ValidateRestrictionResp

	return resp.buildToValidateRestrictionResponse(), nil
}

func (r *ValidateRestrictionResp) buildToValidateRestrictionResponse() restriction2.ValidateRestrictionResponse {
	return restriction2.ValidateRestrictionResponse{}
}
