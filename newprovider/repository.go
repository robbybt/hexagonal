package newprovider

type repositories struct {
	cache      string
	database   string
	httpClient string
	grpcClient string
}

func NewRepositories(a, b, c, d string) DomainRepository {
	return &repositories{
		cache:      a,
		database:   b,
		httpClient: c,
		grpcClient: d,
	}
}
