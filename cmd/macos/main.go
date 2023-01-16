package main

import (
	"github.com/mniak/auto-otp/internal/keysender"
	"github.com/mniak/auto-otp/internal/mock"
	"github.com/samber/lo"
)

func main() {
	sendKeysChan := make(chan string)
	configProvider := mock.NewConfigProvider()
	typingProvider := keysender.New()

	menuEntries := lo.Must(configProvider.GetMenuEntries())

	go func() {
		for {
			code := <-sendKeysChan
			typingProvider.SendKeys(code)
		}
	}()
	showMenu(sendKeysChan, menuEntries)
}
