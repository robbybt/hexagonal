package campaign

type CampaignData struct {
	PaymentProfile            string
	ProductCampaignData       map[int64]*ProductCampaignData
	RedirectPageApplink       string
	RedirectPageUrl           string
	DeductBsOnCanceledPayment bool
}

type ProductCampaignData struct {
	ProductID                int64
	CampaignID               int64
	EventID                  int64
	AppsOnly                 bool
	MaxOrder                 int32
	MinOrder                 int32
	OriginalPrice            int64
	DiscountPrice            int64
	OriginalPriceFmt         string
	DiscountPriceFmt         string
	CustomStock              int32
	CampaignTypeID           int64
	CampaignTypeName         string
	HideGimmick              bool
	BookingStock             int64
	UseWarehouse             bool
	ErrorCode                int32
	RedirectPageApplink      string
	RedirectPageUrl          string
	Applink                  string
	IsBigCampaign            bool
	IsSlashPrice             bool
	IsSlashPriceRemovalQuota bool
	CampaignShortName        string
	StockReference           bool
	CampaignGreaterCityIDs   []int64
	StackingConfig           string
}
