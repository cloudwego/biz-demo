package configparser

import (
	"github.com/spf13/pflag"
)

const (
	configFlagName = "config"
)

var (
	configFlag *string
)

// Flags adds flags related to basic configuration's parser loader to the flags.
func Flags(flags *pflag.FlagSet) {
	configFlag = flags.String(configFlagName, "", "Path to the config file")
}

func getConfigFlag() string {
	return *configFlag
}
