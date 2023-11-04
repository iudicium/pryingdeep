package exporter

import (
	"github.com/spf13/cobra"

	"github.com/pryingbytez/pryingdeep/pkg/cmd/exporter/json"
)

var ExporterCMD = &cobra.Command{
	Use:   "export",
	Short: "Export the collected data into a file.",
	Long:  "Export the collected data from the database into a file.",
}

func init() {
	ExporterCMD.AddCommand(json.JSONCmd)

}
