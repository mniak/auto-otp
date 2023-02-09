package configfile

import (
	"log"
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

func (e ConfigOTPEntry) GetOTP() (chan string, error) {
	codeChan := make(chan string)
	code, err := totp.GenerateCode(e.OTPSecret, time.Now())
	if err != nil {
		return nil, err
	}
	codeChan <- code
	go func() {
		time.Sleep(900 * time.Millisecond)
		code, err := totp.GenerateCode(e.OTPSecret, time.Now())
		if err != nil {
			log.Fatalln("OTP generation failed", err)
		}
		codeChan <- code
	}()
	return codeChan, nil
}
