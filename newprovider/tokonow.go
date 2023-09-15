package newprovider

import (
	"fmt"
	"hexagonal/domain/tokonow"
)

type InputGetTokonow struct {
	CategoryID int
}

func (repo *Repositories) GetTokonowData(categoryID int) (tokonow.TokonowCategory, error) {
	// restapicalls
	fmt.Println("rest API with", categoryID)
	var resp tokonow.TokonowCategory

	return resp, nil
}
