package autootp

type OTPEntries struct {
	Title string
	Code  func() string
}

type OTPEntriesProvider interface {
	OTPEntries() ([]OTPEntries, error)
}
