package autootp

type MenuEntry struct {
	Title string
	Code  func() string
}

type ConfigProvider interface {
	GetMenuEntries() ([]MenuEntry, error)
}
