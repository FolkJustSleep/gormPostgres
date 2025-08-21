package gateway

import (
	"github.com/gofiber/fiber/v2"
	"go-template/src/service"
	"go-template/data/model"
)

type HTTPGateway struct {//for store the service that you want to use
	UserService service.IUserService
	LoginService service.ILoginService
}
func HTTPGatewayHandler(app *fiber.App, userService service.IUserService, loginService service.ILoginService) {
	gateway := &HTTPGateway{
		UserService: userService,
		LoginService: loginService,
	}
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(model.Response{
			Status: 200,
			Message: "This is GO Fiber API ",
			Data: "version 1.0.0",
		})
	})
	gatewayUser(*gateway,app)//sent service you want to use to route
	gatewayLogin(*gateway, app)//sent service you want to use to route
}