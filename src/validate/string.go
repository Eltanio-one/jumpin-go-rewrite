package validate

import (
	"fmt"
	"regexp"
)

const (
	usermailPattern = `^[r"^[a-zA-Z0-9.!#$%&'*+\/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"]`
)

func Usermail(usermail string) error {
	reg := regexp.MustCompile(usermailPattern)

	match := reg.MatchString(usermail)
	if !match {
		return fmt.Errorf("invalid username/email provided")
	}
	return nil
}
