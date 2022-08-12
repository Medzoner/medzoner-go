package captcha

import (
	"github.com/dpapathanasiou/go-recaptcha"
)

//Captcher Captcher
type Captcher interface {
	Confirm(remoteip, response string) (result bool, err error)
}

//RecaptchaAdapter RecaptchaAdapter
type RecaptchaAdapter struct {
}

//Confirm Confirm
func (s RecaptchaAdapter) Confirm(remoteip, response string) (result bool, err error) {
	return recaptcha.Confirm(remoteip, response)
}
