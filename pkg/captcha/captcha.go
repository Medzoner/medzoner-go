package captcha

import (
	"fmt"

	"github.com/dpapathanasiou/go-recaptcha"
)

type Config struct {
	RecaptchaSiteKey   string `env:"SITE_KEY"   envDefault:"xxxxxxxxxxxx"`
	RecaptchaSecretKey string `env:"SECRET_KEY" envDefault:"xxxxxxxxxxxx"`
}

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

func NewRecaptchaAdapter(cfg Config) *RecaptchaAdapter {
	recaptcha.Init(cfg.RecaptchaSecretKey)

	return &RecaptchaAdapter{
		RecaptchaSiteKey:   cfg.RecaptchaSiteKey,
		RecaptchaSecretKey: cfg.RecaptchaSecretKey,
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
