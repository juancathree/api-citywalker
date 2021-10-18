package server

// placeRouter "TravelAPI/pkg/place/infrastructure/router"
// travelRouter "TravelAPI/pkg/travel/infrastructure/router"
import (
	userRouter "citywalker/pkg/user/infrastructure/router"
)

func SetupRouter() {
	//User
	user := app.Group("/")
	userRouter.Router(user)
}
