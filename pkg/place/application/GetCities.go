package application

func GetCities(lng *string) (*[]string, error) {
	return repo.GetCities(lng)
}
