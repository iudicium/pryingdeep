package config

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	criteria           map[string]string
	exporterConfigPath = "configs/json/exporterConfig.json"
	associations       = "all"
	sortBy             = ""
	sortOrder          = "asc"
	limit              = 0
	filePath           = "example.json"
)
var ExporterCMD = &cobra.Command{
	Use:   "exporter",
	Short: "Configure exporterConfig.json from cmd",
	Long:  "Easily configure  exporterConfig.json from the command line",
	Run:   configureExporter,
}

func init() {
	ExporterCMD.Flags().StringToStringVarP(&criteria, "criteria", "c", criteria, "JSON-formatted criteria, e.g., -c 'title=test,\"url=LIKE example.com\"'")

	ExporterCMD.Flags().StringVarP(&exporterConfigPath, "file", "e", exporterConfigPath, "Configuration file path")
	ExporterCMD.Flags().StringVarP(&associations, "associations", "a", associations, "-a WP,E,P,C")
	ExporterCMD.Flags().StringVarP(&sortBy, "sortBy", "s", sortBy, "SortBy e.g -> -b title")
	ExporterCMD.Flags().StringVarP(&sortOrder, "sortOrder", "o", sortOrder, "SortOrder SortBy e.g -> -o ASC || -b DESC")
	ExporterCMD.Flags().IntVarP(&limit, "limit", "l", limit, "Limit e.g -> -l 100 -> 100 items will be taken from database. Default limit will acquire all results from database")
	ExporterCMD.Flags().StringVarP(&filePath, "filePath", "f", filePath, "FilePath -f myfilepath")
}

func configureExporter(cmd *cobra.Command, args []string) {
	viper.Set("Associations", associations)
	viper.Set("SortBy", sortBy)
	viper.Set("SortOrder", sortOrder)
	viper.Set("Limit", limit)
	viper.Set("FilePath", filePath)
	viper.Set("Criteria", criteria)

	viper.SetConfigFile(exporterConfigPath)

	if err := viper.WriteConfig(); err != nil {
		log.Fatalf("Error writing configuration file: %s", err)
	}

	fmt.Printf("[+] Configuration saved to %s\n", exporterConfigPath)
}
