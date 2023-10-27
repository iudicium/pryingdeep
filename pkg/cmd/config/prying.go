package config

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configFile string
	wordpress  bool
	crypto     bool
	email      bool
	phone      []string
)
var PryingCMD = &cobra.Command{
	Use:   "prying",
	Short: "Configure pryingConfig.json from cmd",
	Long:  "Easily configure pryingConfig.jsom from command line.",
	Run:   configurePryingTools,
}

func init() {
	PryingCMD.Flags().StringVarP(&configFile, "file", "f", "configs/json/pryingConfig.json", "Configuration file path")
	PryingCMD.Flags().BoolVarP(&wordpress, "wordpress", "w", false, "Enable WordPress support")
	PryingCMD.Flags().BoolVarP(&crypto, "crypto", "c", false, "Enable crypto features")
	PryingCMD.Flags().BoolVarP(&email, "email", "e", false, "Enable email notifications")
	PryingCMD.Flags().StringSliceVarP(&phone, "phone", "p", []string{}, "Phone numbers for notifications")
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
