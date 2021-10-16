package main
import (
	"os"
	"log"
	"net/http"
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

	// Register default handler, which will use my custom urlconf
	http.HandleFunc("/", urlroute)

	// Run the server
	log.Fatal(http.ListenAndServe(":" + port, nil))

}
