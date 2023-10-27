package cmd

//type exportOptions struct {
//	QueryBuilder querybuilder.QueryBuilder
//}
//
//var exportCmd = &cobra.Command{
//	Use:   "export",
//	Short: "Export the crawled data to json",
//	Long:  longExplanation(),
//	Run:   parseExportArgs,
//}
//
//func init() {
//	options := &exportOptions{}
//
//	exportCmd.Flags().StringSliceVarP(&options.QueryBuilder.WebPageCriteria, "criteria", "c", []string{}, "WebPageCriteria (key=value, key like value, or key without a value)")
//	exportCmd.Flags().StringVarP(&options.QueryBuilder.Associations, "associations", "a", "", "Associations")
//	exportCmd.Flags().StringVarP(&options.QueryBuilder.SortBy, "sort-by", "b", "", "Sort by field")
//	exportCmd.Flags().StringVarP(&options.QueryBuilder.SortOrder, "sort-order", "o", "", "Sort order (e.g., ASC, DESC)")
//	exportCmd.Flags().IntVarP(&options.QueryBuilder.Limit, "limit", "l", 0, "Number of rows to export (0 for all)")
//
//	rootCmd.AddCommand(exportCmd)
//}
//
//func parseExportArgs(cmd *cobra.Command, args []string) {
//	options := &exportOptions{}
//
//	fmt.Println("WebPageCriteria:", options.QueryBuilder.WebPageCriteria)
//	fmt.Println("Associations:", options.QueryBuilder.Associations)
//	fmt.Println("Sort By:", options.QueryBuilder.SortBy)
//	fmt.Println("Sort Order:", options.QueryBuilder.SortOrder)
//	fmt.Println("Limit:", options.QueryBuilder.Limit)
//}
//
//func longExplanation() string {
//	return ""
//}
