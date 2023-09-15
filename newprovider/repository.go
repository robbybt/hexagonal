package repository

import "hexagonal/domain"

type Repositories struct {
	cache    Cache
	database Database
}

func NewRepositories(c Cache, d Database) *Repositories {
	return &Repositories{
		cache:    c,
		database: d,
	}
}

func (repo *Repositories) FetchDomain1() domain.Domain1 {
	// For simplicity, returning static data. Ideally, fetch from an external source
	return domain.Domain1{
		Field1: "Data1_1",
		Field2: "Data1_2",
	}
}

func (repo *Repositories) FetchDomain2() domain.Domain2 {
	// For simplicity, returning static data. Ideally, fetch from an external source
	return domain.Domain2{
		Field1: "Data2_1",
		Field2: "Data2_2",
	}
}

func (repo *Repositories) FetchDomain3(dom1Data, dom2Data string) (domain.Domain3, error) {
	cacheKey := dom1Data + ":" + dom2Data

	// Try to fetch from Cache first
	if cachedValue, err := repo.cache.Get(cacheKey); err == nil && cachedValue != nil {
		return *cachedValue, nil
	}

	// If not found in cache or if there's an error, fetch from database
	dbValue, err := repo.database.FetchByCriteria(dom1Data, dom2Data)
	if err != nil {
		return domain.Domain3{}, err
	}

	// Store the fetched value in cache for subsequent requests
	repo.cache.Set(cacheKey, dbValue)

	return dbValue, nil
}

func (repo *Repositories) FetchDomain4(dom3Field string) domain.Domain4 {
	// For simplicity, fetching from Redis based on dom3Field. Implement logic as needed
	cacheValue, err := repo.cache.Get(dom3Field)
	if err != nil {
		return domain.Domain4{}
	}
	return domain.Domain4{
		Field1:    "Data4_1",
		Field2:    "Data4_2",
		Dom3Field: cacheValue.Field1, // Or some other logic to derive Domain4 data from Domain3
	}
}
