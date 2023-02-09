package configfile

import autootp "github.com/mniak/auto-otp"

var _ autootp.OTPEntriesProvider = &otpEntriesProvider{}
