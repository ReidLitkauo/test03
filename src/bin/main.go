package main

import (
	"os"
	"log"
	"github.com/gofiber/fiber/v2"
)

func main() {

	// exe, err := os.Executable()
	// if err != nil {
	// 	panic(err)
	// }
	// ex_path := filepath.Dir(exe)
	// fmt.Println(ex_path)


	app := fiber.New()
	port := os.Getenv("PORT")

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World 👋!")
	// })
	app.Static("/", "/out/rsc/_index/.html")

	log.Fatal(app.Listen(":" + port))
}
