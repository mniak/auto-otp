package mock

import (
	"time"

	autootp "github.com/mniak/auto-otp"
)

type simpleConfig struct{}

func NewConfigProvider() *simpleConfig {
	return &simpleConfig{}
}

func (sc *simpleConfig) GetMenuEntries() ([]autootp.OTPEntries, error) {
	return []autootp.OTPEntries{
		{
			Title: "Demo",
			Code: func() string {
				time.Sleep(1 * time.Second)
				return "123456"
			},
		},
		{
			Title: "Example",
			Code: func() string {
				time.Sleep(1 * time.Second)
				return "777888"
			},
		},
		{
			Title: "Any Site",
			Code: func() string {
				time.Sleep(1 * time.Second)
				return "000000"
			},
		},
	}, nil
}
