package exporter

import (
	"github.com/spf13/cobra"

	"github.com/pryingbytez/pryingdeep/pkg/cmd/exporter/json"
)

var ExporterCMD = &cobra.Command{
	Use:   "export",
	Short: "Configure exporterConfig.json from cmd",
	Long:  "Export the collected data from the database into a file.",
}

func init() {
	ExporterCMD.AddCommand(json.JSONCmd)

}
