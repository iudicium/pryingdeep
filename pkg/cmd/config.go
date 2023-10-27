package cmd

import (
	"github.com/spf13/cobra"

	"github.com/pryingbytez/prying-deep/pkg/cmd/config"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure prying and exporter settings from cmd",
	Long:  "Easily configure prying and exporter settings from cmd. We do not plan onto adding support for crawlerConfig.json",
}

func init() {
	configCmd.AddCommand(config.PryingCMD)
	configCmd.AddCommand(config.ExporterCMD)
	rootCmd.AddCommand(configCmd)

}
