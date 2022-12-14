package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-server-jwt/controllers"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")

	AuthApi(api)
}

func AuthApi(api fiber.Router) {
	api.Post("/register", controllers.Register)
	api.Post("/login", controllers.Login)
	api.Get("/user", controllers.GetUserWithToken)
	api.Post("/logout", controllers.LogoutWithCookie)
	api.Post("/update_password", controllers.UpdatePassword)
	api.Post("/update_user_info", controllers.UpdateUserInfo)
	api.Post("/delete_user", controllers.DeleteUser)
}
