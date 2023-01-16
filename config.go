package autootp

type OTPEntry struct {
	Title string
	Code  func() (string, error)
}

type OTPEntriesProvider interface {
	OTPEntries() ([]OTPEntry, error)
}
