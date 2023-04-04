package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"gopkg.in/DataDog/dd-trace-go.v1/profiler"
	"log"
)

func main() {

	tracer.Start()
	defer tracer.Stop()

	err := profiler.Start(
		profiler.WithProfileTypes(
			profiler.CPUProfile,
			profiler.HeapProfile,

			// The profiles below are disabled by
			// default to keep overhead low, but
			// can be enabled as needed.
			profiler.BlockProfile,
			profiler.MutexProfile,
			profiler.GoroutineProfile,
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	defer profiler.Stop()

	app := fiber.New()

	app.Get("/server/http", GetPost)

	// Start server on port 3000
	app.Listen(":3001")
}

func GetPost(c *fiber.Ctx) error {
	//200 code - v1
	fmt.Println("----- v1 ---- ")
	fmt.Println("return 200")
	return c.JSON(fiber.Map{"response": "return from v1 - 200"})

	// 500 Service Unavailable - v2
	//fmt.Println("----- v2 ---- ")
	//fmt.Println("return 500")
	//return fiber.ErrInternalServerError

	// We got error in v2 and send request in v3
	//fmt.Println("----- v3 ---- ")
	//fmt.Println("return 200 from v3")
	//return c.JSON(fiber.Map{"error": "return 200 from v3"})

	//fmt.Println("---- Mirroring ----")
	//return c.JSON(fiber.Map{"INFO": "INFO"})

}
