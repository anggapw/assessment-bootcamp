package routes

import (
	"assessment/layer/user"
)

var (
	userService    = user.NewService(userRepository)
	userRepository = user.NewRepository(DB)
	userHandler    = handler.NewUserHandler(userService, authService)
)
