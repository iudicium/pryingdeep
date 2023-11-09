package configs

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type ConfigSaver interface {
	ConfigureViper(*pflag.Flag)
	SaveConfig()
}

// CLIConfig is a struct to store configuration parameters.
type CLIConfig struct {
	Silent     bool
	SaveConfig bool
	ConfigPath string
}

// NewCLIConfig creates a new CLIConfig instance.
func NewCLIConfig() *CLIConfig {
	return &CLIConfig{
		Silent:     viper.GetBool("silent"),
		SaveConfig: viper.GetBool("save-config"),
		ConfigPath: viper.GetString("config"),
	}
}

func (c *CLIConfig) StoreConfig() {
	Save(c.ConfigPath, "config", "save-config", "silent", "name")
	fmt.Println(color.GreenString("[+]"), "Config has been successfully modified at:", color.RedString(c.ConfigPath))
}

// ConfigureViper configures Viper flags for a given prefix.
func (c *CLIConfig) ConfigureViper(prefix string) func(*pflag.Flag) {
	return func(flag *pflag.Flag) {
		prefixedName := strings.Join([]string{prefix, flag.Name}, ".")
		if err := viper.BindPFlag(prefixedName, flag); err != nil {
			panic(err)
		}
		env := strings.ReplaceAll(prefixedName, ".", "_")
		viper.BindEnv(prefixedName, strings.ToUpper(env))
	}
}
