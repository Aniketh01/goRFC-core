package parser

import (
	"net/url"
	"log"
	"strings"
)

// Returns the type of url it is parsing, if not found returns empty string(Need to check of this while using this as an API.)
func getType(str string) string {
	u, err := url.Parse(str)
	if err != nil {
		log.Fatal(err)
	}

	if (strings.Contains(u.Path, "/html/")) {
		return "html"
	} else if (strings.Contains(u.Path, "/pdf/")) {
		return "pdf"
	} else if (strings.Contains(u.Path, "/xml/")) {
		return "xml"
	} else {
		return ""
	}
}