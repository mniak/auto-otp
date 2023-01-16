package keysender

import (
	"fmt"
)

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#include <ApplicationServices/ApplicationServices.h>
*/
import "C"

type keySender struct{}

func New() (*keySender, error) {
	return &keySender{}, nil
}

func (ks *keySender) SendKeys(keys string) error {
	fmt.Println("sending keys", keys)
	for _, k := range keys {
		if err := pressKey(k); err != nil {
			return err
		}
	}
	return nil
}

var keyCodes = map[rune]int{
	'1': 18,
	'2': 19,
	'3': 20,
	'4': 21,
	'5': 23,
	'6': 22,
	'7': 26,
	'8': 28,
	'9': 25,
	'0': 29,
}

func pressKey(key rune) error {
	keyCode := keyCodes[key]
	event := C.CGEventCreateKeyboardEvent(C.CGEventSourceRef(0), C.CGKeyCode(keyCode), true)
	C.CGEventPost(C.kCGSessionEventTap, event)
	return nil
}
