package newprovider

type Repositories struct {
	cache      string
	database   string
	httpClient string
	grpcClient string
}

func NewRepositories(a, b, c, d string) *Repositories {
	return &Repositories{
		cache:      a,
		database:   b,
		httpClient: c,
		grpcClient: d,
	}
}
