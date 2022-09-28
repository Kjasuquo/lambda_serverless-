package validators

import (
	"log"
	"regexp"
)

func IsEmailValid(email string) bool {
	rxEmail := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if len(email) < 3 || len(email) > 254 || !rxEmail.MatchString(email) {
		log.Println("email not valid")
		return false
	}

	return true
}
