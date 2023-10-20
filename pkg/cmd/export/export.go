package export

import (
	"fmt"
	"github.com/r00tk3y/prying-deep/pkg/cmd"
	"github.com/spf13/cobra"
)

type exportOptions struct {
}

var exportCmd = &cobra.Command{
	Use:    "export",
	Short:  "Export the crawled data to json",
	Args:   cobra.MinimumNArgs(2),
	PreRun: parseExportArgs,
}

func init() {
	exportCmd.Flags().StringP("model", "m", "", "Database model name")
	exportCmd.Flags().IntP("rows", "r", 0, "Number of rows to export (0 for all)")
	exportCmd.Flags().StringP("condition", "c", "", "Export condition (optional)")
	exportCmd.Flags().StringP("filename", "f", "exported_data.json", "Output filename (default: exported_data.json)")
	exportCmd.MarkFlagRequired("filename")
	cmd.RootCmd.AddCommand(exportCmd)
}

func parseExportArgs(cmd *cobra.Command, args []string) {
	modelName, _ := cmd.Flags().GetString("model")
	rows, _ := cmd.Flags().GetInt("rows")
	condition, _ := cmd.Flags().GetString("condition")
	filename, _ := cmd.Flags().GetString("filename")
	fmt.Println(modelName, rows, condition, filename)
}
