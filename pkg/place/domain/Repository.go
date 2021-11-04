package domain

type Repository interface {
	Gets(city *string, lng *string) (*[]Place, error)
	GetsCustom(city *string, pref, cuV *[]string, lng *string) (*[]Place, error)
	GetCities(lng *string) (*[]string, error)
}
