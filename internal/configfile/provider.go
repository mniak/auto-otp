package configfile

import (
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	autootp "github.com/mniak/auto-otp"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/spf13/viper"
)

func NewOTPEntiresProvider() (*otpEntriesProvider, error) {
	if homedir, err := os.UserHomeDir(); err == nil {
		viper.AddConfigPath(homedir)
	}
	viper.SetConfigName(".autootp")
	viper.SetConfigType("yaml")

	viper.OnConfigChange(func(in fsnotify.Event) {
		err := viper.Unmarshal(&config)
		log.Println(errors.WithMessage(err, "failed to load config"))
	})
	viper.WatchConfig()

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.WithMessage(err, "failed to load config")
	}
	if err := viper.Unmarshal(&config); err != nil {
		return nil, errors.WithMessage(err, "failed to load config")
	}

	return &otpEntriesProvider{}, nil
}

type otpEntriesProvider struct{}

var config ConfigFile

func (p otpEntriesProvider) OTPEntries() ([]autootp.OTPEntry, error) {
	return lo.Map(config.Entries, func(item ConfigOTPEntry, index int) autootp.OTPEntry {
		return autootp.OTPEntry{
			Title: item.Title,
			Code:  item.GetOTP,
		}
	}), nil
}
