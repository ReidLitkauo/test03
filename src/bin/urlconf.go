package main
import (
	_ "log"
	"github.com/gofiber/fiber/v2"
)

//##############################################################################
func urlconf (fiberapp *fiber.App) {
// Configures all URLs

	// # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # #
	// Publically-visible URLs (don't need an account)

	// Index
	fiberapp.Static("/", "/out/rsc/_index.html")

}