package main

import (
	"time"

	autootp "github.com/mniak/auto-otp"
	"github.com/mniak/auto-otp/internal/keysender"
)

func main() {
	sendKeysChan := make(chan string)

	keySender := keysender.New()
	go func() {
		for {
			code := <-sendKeysChan
			keySender.SendKeys(code)
		}
	}()
	showMenu(sendKeysChan, []autootp.MenuEntry{
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
	})
}
