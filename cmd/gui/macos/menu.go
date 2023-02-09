package main

import (
	"fmt"
	"log"
	"regexp"

	"github.com/getlantern/systray"
	autootp "github.com/mniak/auto-otp"
)

func showMenu(sendKeysChan chan<- string, menuEntries []autootp.OTPEntry) {
	systray.Run(
		func() { // tray init
			initMenu(sendKeysChan, menuEntries)
		},
		func() { // tray exit
		},
	)
}

func addMenuItem(entry autootp.OTPEntry, sendKeysChan chan<- string) *systray.MenuItem {
	menuItem := systray.AddMenuItem(fmt.Sprintf("%s - Loading...", entry.Title), "")
	menuItem.Disable()
	codeChan, err := entry.Code()
	if err != nil {
		menuItem.SetTitle(fmt.Sprintf("%s - %s", entry.Title, "Error!"))
		menuItem.Disable()
		log.Println(err)
		return menuItem
	}
	var currentCode string
	go func() {
		for {
			currentCode = <-codeChan
			menuItem.SetTitle(fmt.Sprintf("%s - %s", entry.Title, prettyCode(currentCode)))
			menuItem.Enable()
		}
	}()
	go func() {
		for {
			<-menuItem.ClickedCh
			sendKeysChan <- currentCode
		}
	}()
	return menuItem
}

func initMenu(sendKeysChan chan<- string, entries []autootp.OTPEntry) {
	systray.SetTitle("Auto OTP")
	subtitle := systray.AddMenuItem("Click to type OTP", "")
	subtitle.Disable()
	for _, entry := range entries {
		addMenuItem(entry, sendKeysChan)
	}
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "")
	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()
}

var rePrettify = regexp.MustCompile(`(\d{3})(\d)`)

func prettyCode(code string) string {
	return rePrettify.ReplaceAllString(code, "$1-$2")
}
