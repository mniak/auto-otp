package main

import (
	"fmt"
	"time"

	"github.com/getlantern/systray"
)

type Entry struct {
	Title   string
	GetCode func() string
}

func main() {
	onReady := func() {
		systray.SetTitle("Auto OTP")

		entries := []Entry{
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
		}
		for _, entry := range entries {
			menuItem := systray.AddMenuItem(fmt.Sprintf("%s - Loading...", entry.Title), "")
			menuItem.Disable()
			go func(mi *systray.MenuItem, e Entry) {
				time.Sleep(1 * time.Second)
				code := e.GetCode()
				mi.SetTitle(fmt.Sprintf("%s - %s", e.Title, code))
				mi.Enable()
			}(menuItem, entry)
		}

		systray.AddSeparator()

		mQuit := systray.AddMenuItem("Quit", "")
		go func() {
			<-mQuit.ClickedCh
			systray.Quit()
		}()
	}
	onExit := func() {
	}
	systray.Run(onReady, onExit)
}
