package main

import "fmt"

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
			Title:   "Demo",
			GetCode: func() string { return "123-456" },
		},
		{
			Title:   "Example",
			GetCode: func() string { return "777-888" },
		},
		{
			Title:   "Any Site",
			GetCode: func() string { return "000-000" },
		},
	})
}
