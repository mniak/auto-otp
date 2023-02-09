package autootp

type OTPEntry struct {
	Title string
	Code  func() (chan string, error)
}

type OTPEntriesProvider interface {
	OTPEntries() ([]OTPEntry, error)
}
