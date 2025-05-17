package email

import (
	"strings"

	"github.com/pass1on-ok/EmailVerification/pkg/helpers"
)

func VerifyEmail(email string) error {

	if err := helpers.ValidateEmail(email); err != nil {
		return err
	}

	domain := email[strings.Index(email, "@")+1:]
	if err := CheckDomain(domain); err != nil {
		return err
	}

	return nil
}
