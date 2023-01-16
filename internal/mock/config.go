package mock

import (
	"time"

	autootp "github.com/mniak/auto-otp"
)

type simpleConfig struct{}

func NewConfigProvider() *simpleConfig {
	return &simpleConfig{}
}

func (sc *simpleConfig) GetMenuEntries() ([]autootp.MenuEntry, error) {
	return []autootp.MenuEntry{
		{
			Title: "Demo",
			Code: func() string {
				time.Sleep(1 * time.Second)
				return "123-456"
			},
		},
		{
			Title: "Example",
			Code: func() string {
				time.Sleep(1 * time.Second)
				return "777-888"
			},
		},
		{
			Title: "Any Site",
			Code: func() string {
				time.Sleep(1 * time.Second)
				return "000-000"
			},
		},
	}, nil
}
