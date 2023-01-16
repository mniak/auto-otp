package main

import (
	"fmt"
	"time"
)

func main() {
	sendKeysChan := make(chan string)
	go func() {
		for {
			fmt.Println("waiting for keys")
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
