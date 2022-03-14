package api

import (
	"strings"

	"github.com/asaskevich/govalidator"
)

func validateURL(url string) bool {
	return govalidator.IsURL(url)
}

func enforceHTTP(url string) string {
	if strings.HasPrefix(url, "http") {
		return url
	}

	return "http://" + url
}
