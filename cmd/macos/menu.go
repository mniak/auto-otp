package main

import (
	"fmt"
	"time"

	"github.com/getlantern/systray"
)

type MenuEntry struct {
	Title   string
	GetCode func() string
}

func showMenu(sendKeysChan chan<- string, menuEntries []MenuEntry) {
	systray.Run(
		func() { // tray init
			initMenu(sendKeysChan, menuEntries)
		},
		func() { // tray exit
		},
	)
}

func initMenu(sendKeysChan chan<- string, entries []MenuEntry) {
	systray.SetTitle("Auto OTP")
	subtitle := systray.AddMenuItem("Click to type OTP", "")
	subtitle.Disable()
	for _, entry := range entries {
		menuItem := systray.AddMenuItem(fmt.Sprintf("%s - Loading...", entry.Title), "")
		menuItem.Disable()
		go func(mi *systray.MenuItem, e MenuEntry) {
			time.Sleep(1 * time.Second)
			code := e.GetCode()
			mi.SetTitle(fmt.Sprintf("%s - %s", e.Title, code))
			mi.Enable()
			for {
				<-mi.ClickedCh
				sendKeysChan <- code
			}
		}(menuItem, entry)
	}
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "")
	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()
}
