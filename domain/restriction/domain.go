package restriction

type ValidateRestrictionResponse struct {
	Success      bool
	Message      string
	DataResponse []DataResponse
	MetaResponse []MetaResponse
}

type DataResponse struct {
	ProductId int64
	Status    string
	Action    []string
}

type MetaResponse struct {
	RestrictionName string
	DataResponse    []DataResponse
	ErrorMessage    []string
}
