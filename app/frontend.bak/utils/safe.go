package utils

import "net/url"

var validHost = []string{
	"localhost:8080",
}

func ValidateNext(next string) bool {
	urlObj, err := url.Parse(next)
	if err != nil {
		return false
	}
	if InArray(urlObj.Host, validHost) {
		return true
	}
	return false
}
