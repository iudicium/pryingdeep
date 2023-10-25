package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

type exportOptions struct {
}

var exportCmd = &cobra.Command{
	Use:   "exporters",
	Short: "Export the crawled data to json",
	Run:   parseExportArgs,
}

func init() {
	exportCmd.Flags().StringP("model", "m", "", "Database model name")
	exportCmd.Flags().IntP("rows", "r", 0, "Number of rows to exporters (0 for all)")
	exportCmd.Flags().StringP("condition", "c", "", "Export condition (optional)")
	exportCmd.Flags().StringP("filename", "f", "exported_data.json", "Output filename (default: exported_data.json)")
	exportCmd.MarkFlagRequired("rows")
	rootCmd.AddCommand(exportCmd)
}

func parseExportArgs(cmd *cobra.Command, args []string) {
	modelName, _ := cmd.Flags().GetString("model")
	rows, _ := cmd.Flags().GetInt("rows")
	condition, _ := cmd.Flags().GetString("condition")
	filename, _ := cmd.Flags().GetString("filename")
	fmt.Println(modelName, rows, condition, filename)
}
