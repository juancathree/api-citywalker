package server

import (
	placeRouter "citywalker/pkg/place/infrastructure/router"
	travelRouter "citywalker/pkg/travel/infrastructure/router"
	"citywalker/pkg/user/infrastructure/middleware"
	userRouter "citywalker/pkg/user/infrastructure/router"
)

func SetupRouter() {
	// User
	user := app.Group("/")
	userRouter.Router(user)

	// Place
	place := app.Group("/:lng/place", middleware.AuthRequired())
	placeRouter.Router(place)

	// Travel
	travel := app.Group("/:lng/travel", middleware.AuthRequired())
	travelRouter.Router(travel)
}
