package tokonow

type TokonowRepository interface {
	GetTokonowData(categoryID int) (TokonowCategory, error)
}
