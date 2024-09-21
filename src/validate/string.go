package validate

import (
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

const (
	usermailPattern       = `^(?:[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}|[a-zA-Z0-9._%+-]{3,20})$`
	usernamePattern       = `^[a-zA-Z0-9._%+-]{3,20})$`
	emailPattern          = `^(?:[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,})$`
	minimumIntegerPattern = `[0-9]`
	minimumSpecialPattern = `[!@#\$%\^&\*\(\)_\+\[\]\{\}\|;:'",.<>?/\\]`
)

func Username(username string) error {
	reg := regexp.MustCompile(usernamePattern)

	match := reg.MatchString(username)
	if !match {
		return fmt.Errorf("invalid username provided")
	}
	return nil
}

func Email(email string) error {
	reg := regexp.MustCompile(emailPattern)

	match := reg.MatchString(email)
	if !match {
		return fmt.Errorf("invalid email provided")
	}
	return nil
}

func UsernameOrEmail(email string) bool {
	reg := regexp.MustCompile(emailPattern)

	return reg.MatchString(email)
}

// password must contain:
// min 8 characters
// min 1 number, min 1 special
func Password(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters in length")
	}

	reg := regexp.MustCompile(minimumIntegerPattern)
	containsInt := reg.MatchString(password)
	if !containsInt {
		return fmt.Errorf("password must contain at least 1 number")
	}

	reg = regexp.MustCompile(minimumSpecialPattern)
	containsSpecial := reg.MatchString(password)
	if !containsSpecial {
		return fmt.Errorf("password must contain at least 1 special character")
	}

	return nil
}

func Hash(hash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
