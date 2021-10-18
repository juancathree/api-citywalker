package application

func PostTravel(userID *string, travelID *string) error {
	return repo.PostTravel(userID, travelID)
}
