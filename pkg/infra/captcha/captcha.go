package captcha

import (
	"fmt"

	"github.com/dpapathanasiou/go-recaptcha"
)

type RecaptchaSiteKey string

type Captcher interface {
	Confirm(remoteip, response string) (result bool, err error)
}

type RecaptchaAdapter struct{}

func NewRecaptchaAdapter() *RecaptchaAdapter {
	return &RecaptchaAdapter{}
}

func (s RecaptchaAdapter) Confirm(remoteip, response string) (result bool, err error) {
	r, err := recaptcha.Confirm(remoteip, response)
	if err != nil {
		return false, fmt.Errorf("captcha server error: %w", err)
	}
	return r, nil
}
