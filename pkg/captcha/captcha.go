package captcha

import (
	"fmt"

	"github.com/Medzoner/medzoner-go/internal/config"
	"github.com/dpapathanasiou/go-recaptcha"
)

type RecaptchaSiteKey string

type Captcher interface {
	Confirm(remoteip, response string) (result bool, err error)
	GetSiteKey() string
	GetSecretKey() string
}

type RecaptchaAdapter struct {
	RecaptchaSiteKey   string
	RecaptchaSecretKey string
}

func NewRecaptchaAdapter(cfg config.Config) *RecaptchaAdapter {
	recaptcha.Init(cfg.Recaptcha.RecaptchaSecretKey)

	return &RecaptchaAdapter{
		RecaptchaSiteKey:   cfg.Recaptcha.RecaptchaSiteKey,
		RecaptchaSecretKey: cfg.Recaptcha.RecaptchaSecretKey,
	}
}

func (s RecaptchaAdapter) Confirm(remoteip, response string) (result bool, err error) {
	r, err := recaptcha.Confirm(remoteip, response)
	if err != nil {
		return false, fmt.Errorf("captcha server error: %w", err)
	}
	return r, nil
}

func (s RecaptchaAdapter) GetSiteKey() string {
	return s.RecaptchaSiteKey
}

func (s RecaptchaAdapter) GetSecretKey() string {
	return s.RecaptchaSecretKey
}
