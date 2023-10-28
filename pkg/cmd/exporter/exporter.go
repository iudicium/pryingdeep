package exporter

import (
	"github.com/spf13/cobra"

	"github.com/pryingbytez/prying-deep/pkg/cmd/exporter/json"
)

var ExporterCMD = &cobra.Command{
	Use:   "exporter",
	Short: "Configure exporterConfig.json from cmd",
	Long:  "Export the collected data from the database into a file.",
}

func init() {
	ExporterCMD.AddCommand(json.JSONCmd)

}
