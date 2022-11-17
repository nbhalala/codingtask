/*
 * Verifications : helpers
 * Here we cad add all different kinds of verifications (e.g. Invalid Domain)
 */

package helpers

import (
	"os"
	"strings"
)

// Verify the URL
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
