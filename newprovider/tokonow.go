package newprovider

import (
	"fmt"
	"hexagonal/domain/entities/tokonow"
)

func (repo *repositories) GetTokonowData(categoryID int) (tokonow.TokonowCategory, error) {
	// restapicalls
	fmt.Println("rest API with", categoryID)
	var resp tokonow.TokonowCategory

	return resp, nil
}
