package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var (
	configFile string
	wordpress  bool
	crypto     bool
	email      bool
	phone      []string
)
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure application settings",
	Run:   configurePryingTools,
}

func init() {
	RootCmd.AddCommand(configCmd)
	configCmd.Flags().StringVarP(&configFile, "file", "f", "pryingConfig.json", "Configuration file path")
	configCmd.Flags().BoolVarP(&wordpress, "wordpress", "w", false, "Enable WordPress support")
	configCmd.Flags().BoolVarP(&crypto, "crypto", "c", false, "Enable crypto features")
	configCmd.Flags().BoolVarP(&email, "email", "e", false, "Enable email notifications")
	configCmd.Flags().StringSliceVarP(&phone, "phone", "p", []string{}, "Phone numbers for notifications")
}
func configurePryingTools(cmd *cobra.Command, args []string) {
	viper.Set("WordPress", wordpress)
	viper.Set("Crypto", crypto)
	viper.Set("Email", email)
	viper.Set("PhoneNumbers", phone)

	viper.SetConfigFile(configFile)

	if err := viper.WriteConfig(); err != nil {
		log.Fatalf("Error writing configuration file: %s", err)
	}

	fmt.Printf("[+] Configuration saved to %s\n", configFile)
}
