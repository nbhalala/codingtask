package helpers

import(
	"os"
	"strings"
)

func RemoveDomainError(url string) bool {

	if url == os.Getenv("APP_DOMAIN"){
		return false
	}

	newURL := strings.Replace(url, "http://", "", 1)
	newURL = strings.Replace(newURL, "https://", "", 1)
	newURL = strings.Replace(newURL, "www.", "", 1)
	newURL = strings.Split(newURL, "/")[0]
	
	if newURL == os.Getenv("APP_DOMAIN"){
		return false
	}

	return true
}