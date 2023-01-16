package main

import (
	"time"

	"github.com/mniak/auto-otp/internal/macos"
	"github.com/mniak/auto-otp/internal/mock"
	"github.com/samber/lo"
)

func main() {
	sendKeysChan := make(chan string)
	configProvider := mock.NewConfigProvider()
	typingProvider := lo.Must(macos.New())

	menuEntries := lo.Must(configProvider.GetMenuEntries())
	go func() {
		for {
			code := <-sendKeysChan
			time.Sleep(20 * time.Millisecond)
			typingProvider.SendKeys(code)
		}
	}()
	showMenu(sendKeysChan, menuEntries)
}
