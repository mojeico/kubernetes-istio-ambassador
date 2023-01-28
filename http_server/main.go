package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/server/http", GetPost)

	// Start server on port 3000
	app.Listen(":3001")
}

func GetPost(c *fiber.Ctx) error {
	//200 code - v1
	//fmt.Println("----- v1 ---- ")
	//fmt.Println("return 200")
	//return c.JSON(fiber.Map{"response": "return from v1 - 200"})

	a := c.GetReqHeaders()
	println(a)
	// 500 Service Unavailable - v2
	fmt.Println("----- v2 ---- ")
	fmt.Println("return 500")
	return fiber.ErrInternalServerError

	// We got error in v2 and send request in v3
	//fmt.Println("----- v3 ---- ")
	//fmt.Println("return 200 from v3")
	//return c.JSON(fiber.Map{"error": "return 200 from v3"})

}
