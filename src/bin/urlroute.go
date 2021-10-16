package main
import (
	"net/http"
	_ "log"
	"strings"
)

//##############################################################################
func urlroute (rsp http.ResponseWriter, req *http.Request) {
// Route all URLs according to the urlconf

	// Loop through the entire urlconf
	for k, v := range urlconf {

		// Attempt to match the URL to the regex in this entry's key
		match := k.FindAllStringSubmatch(strings.ToLower(req.URL.Path), -1)

		// If we have a match, call the corresponding handler function
		// and stop searching for additional matches
		if match != nil {
			v(rsp, req, match[0])
			break
		}

	}

}
