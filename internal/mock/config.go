package mock

import (
	autootp "github.com/mniak/auto-otp"
)

type simpleConfig struct{}

func NewConfigProvider() *simpleConfig {
	return &simpleConfig{}
}

func fakeCodeGenerator(code string) func() (chan string, error) {
	return func() (chan string, error) {
		codeChan := make(chan string)
		codeChan <- <-codeChan
		return codeChan, nil
	}
}

func (sc *simpleConfig) GetMenuEntries() ([]autootp.OTPEntry, error) {
	return []autootp.OTPEntry{
		{
			Title: "Demo",
			Code:  fakeCodeGenerator("123456"),
		},
		{
			Title: "Example",
			Code:  fakeCodeGenerator("777888"),
		},
		{
			Title: "Any Site",
			Code:  fakeCodeGenerator("000000"),
		},
	}, nil
}
