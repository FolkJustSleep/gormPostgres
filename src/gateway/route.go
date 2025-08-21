package gateway

import (
	"github.com/gofiber/fiber/v2"
	"go-template/util"
	"go-template/src/middleware"
)

func gatewayUser(gateway HTTPGateway, app *fiber.App){
	api := app.Group("/api/user")
	api.Get("/getall", middleware.CheckRole, gateway.GetAllUser)
	api.Post("/create", gateway.CreateUser)
	api.Get("/get", gateway.GetUserByID)
	api.Put("/update", gateway.UpdateUser)
	api.Delete("/delete", gateway.DeleteUser)
	api.Get("/ip", util.GetIP)
}

func gatewayLogin(gateway HTTPGateway, app *fiber.App) {
	api := app.Group("/api/login")
	api.Post("/login", gateway.Login)
}