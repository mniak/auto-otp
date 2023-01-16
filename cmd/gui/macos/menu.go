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

func initMenu(sendKeysChan chan<- string, entries []autootp.OTPEntry) {
	systray.SetTitle("Auto OTP")
	subtitle := systray.AddMenuItem("Click to type OTP", "")
	subtitle.Disable()
	for _, entry := range entries {
		menuItem := systray.AddMenuItem(fmt.Sprintf("%s - Loading...", entry.Title), "")
		menuItem.Disable()
		go func(mi *systray.MenuItem, e autootp.OTPEntry) {
			code, err := e.Code()
			if err != nil {
				mi.SetTitle(fmt.Sprintf("%s - %s", e.Title, "Error!"))
				mi.Disable()
				log.Println(err)
				return
			}
			mi.SetTitle(fmt.Sprintf("%s - %s", e.Title, prettyCode(code)))
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

var rePrettify = regexp.MustCompile(`(\d{3})(\d)`)

func prettyCode(code string) string {
	return rePrettify.ReplaceAllString(code, "$1-$2")
}
