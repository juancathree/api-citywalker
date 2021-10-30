package domain

type Repository interface {
	Get(city *string) (*MatrixCost, error)
}
