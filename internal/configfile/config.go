package configfile

import (
	"time"

	"github.com/pquerna/otp/totp"
)

type ConfigFile struct {
	Entries []ConfigOTPEntry `mapstructure:"entries"`
}

type ConfigOTPEntry struct {
	Title     string `mapstructure:"title"`
	OTPSecret string `mapstructure:"secret"`
}

func (e ConfigOTPEntry) GetOTP() (string, error) {
	return totp.GenerateCode(e.OTPSecret, time.Now())
}
