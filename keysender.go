package autootp

type KeySender interface {
	SendKeys(keys string, enter bool) error
}
