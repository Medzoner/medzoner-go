package captcha

import (
	"github.com/dpapathanasiou/go-recaptcha"
)

type RecaptchaSiteKey string

type Captcher interface {
	Confirm(remoteip, response string) (result bool, err error)
}

type RecaptchaAdapter struct {
}

func NewRecaptchaAdapter() *RecaptchaAdapter {
	return &RecaptchaAdapter{}
}

func (s RecaptchaAdapter) Confirm(remoteip, response string) (result bool, err error) {
	return recaptcha.Confirm(remoteip, response)
}
