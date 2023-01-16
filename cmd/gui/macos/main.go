package main

import (
	"time"

	"github.com/mniak/auto-otp/internal/configfile"
	"github.com/mniak/auto-otp/internal/macos"
	"github.com/samber/lo"
)

func main() {
	sendKeysChan := make(chan string)
	entriesProvider := lo.Must(configfile.NewOTPEntiresProvider())
	typingProvider := lo.Must(macos.New())

	menuEntries := lo.Must(entriesProvider.OTPEntries())
	go func() {
		for {
			code := <-sendKeysChan
			time.Sleep(20 * time.Millisecond)
			typingProvider.SendKeys(code, true)
		}
	}()
	showMenu(sendKeysChan, menuEntries)
}
