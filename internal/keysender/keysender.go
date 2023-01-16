package keysender

import "fmt"

type keySender struct{}

func New() *keySender {
	return &keySender{}
}

func (ks *keySender) SendKeys(keys string) error {
	fmt.Println("sending keys", keys)
	return nil
}
