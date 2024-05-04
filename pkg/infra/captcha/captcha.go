package captcha

import (
	"github.com/dpapathanasiou/go-recaptcha"
)

type RecaptchaSiteKey string

// Captcher Captcher
type Captcher interface {
	Confirm(remoteip, response string) (result bool, err error)
}

// RecaptchaAdapter RecaptchaAdapter
type RecaptchaAdapter struct {
}

// NewRecaptchaAdapter NewRecaptchaAdapter
func NewRecaptchaAdapter() *RecaptchaAdapter {
	return &RecaptchaAdapter{}
}

// Confirm Confirm
func (s RecaptchaAdapter) Confirm(remoteip, response string) (result bool, err error) {
	return recaptcha.Confirm(remoteip, response)
}
