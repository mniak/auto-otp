package autootp

type KeySender interface {
	SendKeys(keys string) error
}
