package newprovider

import (
	"fmt"
	"hexagonal/domain/entities/campaign"
)

type campaignGRPC struct {
	Products                  []*cmpPrdComplete `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	TotalProduct              int64             `protobuf:"varint,2,opt,name=total_product,json=totalProduct,proto3" json:"total_product,omitempty"`
	PaymentProfile            string            `protobuf:"bytes,3,opt,name=payment_profile,json=paymentProfile,proto3" json:"payment_profile,omitempty"`
	DeductBsOnCanceledPayment bool              `protobuf:"varint,5,opt,name=deduct_bs_on_canceled_payment,json=deductBsOnCanceledPayment,proto3" json:"deduct_bs_on_canceled_payment,omitempty"`
	TextDurationTeaser        string            `protobuf:"bytes,6,opt,name=text_duration_teaser,json=textDurationTeaser,proto3" json:"text_duration_teaser,omitempty"`
	CampaignNplNext           string            `protobuf:"bytes,7,opt,name=campaign_npl_next,json=campaignNplNext,proto3" json:"campaign_npl_next,omitempty"`
	HasNextPage               bool              `protobuf:"varint,8,opt,name=has_next_page,json=hasNextPage,proto3" json:"has_next_page,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}          `json:"-"`
	XXX_unrecognized          []byte            `json:"-"`
	XXX_sizecache             int32             `json:"-"`
}

type cmpPrdComplete struct {
	ID                    int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                  string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	URL                   string   `protobuf:"bytes,3,opt,name=URL,proto3" json:"URL,omitempty"`
	URLApps               string   `protobuf:"bytes,4,opt,name=URLApps,proto3" json:"URLApps,omitempty"`
	URLMobile             string   `protobuf:"bytes,5,opt,name=URLMobile,proto3" json:"URLMobile,omitempty"`
	ImageUrl              string   `protobuf:"bytes,6,opt,name=imageUrl,proto3" json:"imageUrl,omitempty"`
	ImageURL700           string   `protobuf:"bytes,7,opt,name=imageURL700,proto3" json:"imageURL700,omitempty"`
	Price                 string   `protobuf:"bytes,8,opt,name=price,proto3" json:"price,omitempty"`
	CourierCount          int64    `protobuf:"varint,11,opt,name=courier_count,json=courierCount,proto3" json:"courier_count,omitempty"`
	Condition             int64    `protobuf:"varint,12,opt,name=condition,proto3" json:"condition,omitempty"`
	DepartmentID          []int64  `protobuf:"varint,13,rep,packed,name=departmentID,proto3" json:"departmentID,omitempty"`
	Rating                float32  `protobuf:"fixed32,16,opt,name=rating,proto3" json:"rating,omitempty"`
	StarRating            float32  `protobuf:"fixed32,17,opt,name=star_rating,json=starRating,proto3" json:"star_rating,omitempty"`
	CountReview           int32    `protobuf:"varint,18,opt,name=count_review,json=countReview,proto3" json:"count_review,omitempty"`
	CountSold             int64    `protobuf:"varint,19,opt,name=count_sold,json=countSold,proto3" json:"count_sold,omitempty"`
	SKU                   string   `protobuf:"bytes,20,opt,name=SKU,proto3" json:"SKU,omitempty"`
	Stock                 int32    `protobuf:"varint,21,opt,name=stock,proto3" json:"stock,omitempty"`
	Returnable            int64    `protobuf:"varint,23,opt,name=returnable,proto3" json:"returnable,omitempty"`
	Status                int32    `protobuf:"varint,24,opt,name=status,proto3" json:"status,omitempty"`
	HasCashback           bool     `protobuf:"varint,25,opt,name=has_cashback,json=hasCashback,proto3" json:"has_cashback,omitempty"`
	CashbackAmount        float32  `protobuf:"fixed32,26,opt,name=cashback_amount,json=cashbackAmount,proto3" json:"cashback_amount,omitempty"`
	MapId                 int64    `protobuf:"varint,27,opt,name=map_id,json=mapId,proto3" json:"map_id,omitempty"`
	LockStatus            int32    `protobuf:"varint,28,opt,name=lock_status,json=lockStatus,proto3" json:"lock_status,omitempty"`
	MaxOrder              int32    `protobuf:"varint,29,opt,name=max_order,json=maxOrder,proto3" json:"max_order,omitempty"`
	PriceUnfmt            string   `protobuf:"bytes,30,opt,name=price_unfmt,json=priceUnfmt,proto3" json:"price_unfmt,omitempty"`
	IsVariant             bool     `protobuf:"varint,31,opt,name=is_variant,json=isVariant,proto3" json:"is_variant,omitempty"`
	ParentId              int64    `protobuf:"varint,32,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	VariantsFilter        []string `protobuf:"bytes,33,rep,name=variants_filter,json=variantsFilter,proto3" json:"variants_filter,omitempty"`
	ChildIds              []int64  `protobuf:"varint,34,rep,packed,name=child_ids,json=childIds,proto3" json:"child_ids,omitempty"`
	SiblingIds            []int64  `protobuf:"varint,35,rep,packed,name=sibling_ids,json=siblingIds,proto3" json:"sibling_ids,omitempty"`
	EggCrackingValidation bool     `protobuf:"varint,36,opt,name=egg_cracking_validation,json=eggCrackingValidation,proto3" json:"egg_cracking_validation,omitempty"`
	MinOrder              int32    `protobuf:"varint,37,opt,name=min_order,json=minOrder,proto3" json:"min_order,omitempty"`
	Position              int64    `protobuf:"varint,40,opt,name=position,proto3" json:"position,omitempty"`
	// errors
	ErrorCode            int32    `protobuf:"varint,41,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	ErrorMessage         string   `protobuf:"bytes,42,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	UseSpoilerPrice      bool     `protobuf:"varint,44,opt,name=use_spoiler_price,json=useSpoilerPrice,proto3" json:"use_spoiler_price,omitempty"`
	IsOutOfCoverage      bool     `protobuf:"varint,45,opt,name=is_out_of_coverage,json=isOutOfCoverage,proto3" json:"is_out_of_coverage,omitempty"`
	ActiveProductCard    bool     `protobuf:"varint,46,opt,name=active_product_card,json=activeProductCard,proto3" json:"active_product_card,omitempty"`
	ShowStockbar         bool     `protobuf:"varint,47,opt,name=show_stockbar,json=showStockbar,proto3" json:"show_stockbar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

type campaignResponse struct {
	campaignGRPC
}

func (repo *repositories) GetCampaign(i campaign.InputGetCampaign) (campaign.CampaignData, error) {
	// restapicalls
	fmt.Println("rest API with", i)

	var resp campaignResponse

	return resp.BuildToCampaignData(), nil
}

func (r *campaignResponse) BuildToCampaignData() campaign.CampaignData {
	return campaign.CampaignData{}
}
