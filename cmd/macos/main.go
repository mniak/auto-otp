package main

import (
	"time"
)

func main() {
	sendKeysChan := make(chan string)
	go func() {
		for {
			code := <-sendKeysChan
			sendKeys(code)
		}
	}()
	showMenu(sendKeysChan, []MenuEntry{
		{
			Title: "Demo",
			GetCode: func() string {
				time.Sleep(1 * time.Second)
				return "123-456"
			},
		},
		{
			Title: "Example",
			GetCode: func() string {
				time.Sleep(1 * time.Second)
				return "777-888"
			},
		},
		{
			Title: "Any Site",
			GetCode: func() string {
				time.Sleep(1 * time.Second)
				return "000-000"
			},
		},
	})
}
