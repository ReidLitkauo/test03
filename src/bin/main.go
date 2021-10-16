package main
import (
	"os"
	"log"
	"github.com/gofiber/fiber/v2"
)

//##############################################################################
func main () {
// Entry point for the entire freaking service

	// # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # #
	// Determine what port to serve on

	// Google App Engine will decide for us what port this service will be on
	// Let's grab that port number as a string from the environment
	port := os.Getenv("PORT")

	// TODO default port for testing only, log.Fatal in production
	if port == "" {
		port = "8080"
	}

	// # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # #
	// Set up HTTP server

	// Create new Fiber app with custom configuration
	fiberapp := fiber.New(fiber.Config{
		AppName: "feedist.app",
	})

	// Configure URLs
	urlconf(fiberapp)

	// Run the server
	log.Fatal(fiberapp.Listen(":" + port))

}
