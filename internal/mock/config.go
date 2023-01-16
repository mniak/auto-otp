package mock

import (
	"time"

	autootp "github.com/mniak/auto-otp"
)

type simpleConfig struct{}

func NewConfigProvider() *simpleConfig {
	return &simpleConfig{}
}

func (sc *simpleConfig) GetMenuEntries() ([]autootp.OTPEntry, error) {
	return []autootp.OTPEntry{
		{
			Title: "Demo",
			Code: func() (string, error) {
				time.Sleep(1 * time.Second)
				return "123456", nil
			},
		},
		{
			Title: "Example",
			Code: func() (string, error) {
				time.Sleep(1 * time.Second)
				return "777888", nil
			},
		},
		{
			Title: "Any Site",
			Code: func() (string, error) {
				time.Sleep(1 * time.Second)
				return "000000", nil
			},
		},
	}, nil
}
