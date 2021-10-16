package main
import (
	"io"
	"net/http"
	"regexp"
	_ "log"
)

//##############################################################################
var urlconf = map[*regexp.Regexp]func(http.ResponseWriter, *http.Request, []string) {
// Bigly config mapping regexes to functions handling them.
// Regexes are the keys, functions are the values.
// If the regex matches the request URL, then function gets called.

	// # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # #
	// Standard pages

	// Index
	regexp.MustCompile("^/$"): myfunc,

	// Legal
	regexp.MustCompile("^/legal/18usc2257$"): myfunc2,

	regexp.MustCompile("^/test/(.*)/(.*)$"): myfunc3,

}

func myfunc (rsp http.ResponseWriter, req *http.Request, args []string) {
	io.WriteString(rsp, "index")
}

func myfunc2 (rsp http.ResponseWriter, req *http.Request, args []string) {
	io.WriteString(rsp, "2257")
}

func myfunc3 (rsp http.ResponseWriter, req *http.Request, args []string) {
	io.WriteString(rsp, "args1: " + args[1] + " - args2: " + args[2])
}