package json

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/iudicium/pryingdeep/configs"
	"github.com/iudicium/pryingdeep/models"
	"github.com/iudicium/pryingdeep/pkg/exporters"
	"github.com/iudicium/pryingdeep/pkg/querybuilder"
)

var JSONCmd = &cobra.Command{
	Use:   "json",
	Short: "Export the crawled data to json",
	RunE: func(cmd *cobra.Command, args []string) error {
		return ExportJSON(cmd, args)
	},
}

var (
	rawSQL       = false
	rawFilePath  = "pkg/querybuilder/queries/select.sql"
	criteria     map[string]string
	associations = "all"
	sortBy       = "status_code"
	sortOrder    = "asc"
	limit        = 0
	filePath     = "data.json"
	offset       = 0
)

func init() {
	JSONCmd.Flags().BoolVarP(&rawSQL, "raw-sql", "r", rawSQL, "--raw to use raw sql queries that you provide. All other flags except silent, rawFilePath and filepath will not matter.")
	JSONCmd.Flags().StringVarP(&rawFilePath, "raw-sql-filepath", "p", rawFilePath, "-rp to specify the file path to the sql file. Only use this flag if you specify -raw")
	JSONCmd.Flags().StringToStringVarP(&criteria, "criteria", "q", criteria, "JSON-formatted criteria, e.g., -q 'title=test,\"url=LIKE example.com\"'")
	JSONCmd.Flags().StringVarP(&associations, "associations", "a", associations, "-a WP,E,P,C")
	JSONCmd.Flags().StringVarP(&sortBy, "sort-by", "b", sortBy, "SortBy e.g -> -b title")
	JSONCmd.Flags().StringVarP(&sortOrder, "sort-order", "", sortOrder, "SortOrder e.g -> -o ASC || -b DESC. Only use this flag if you use SortBy")
	JSONCmd.Flags().IntVarP(&limit, "limit", "l", limit, "Limit e.g -> -l 100 -> 100 items will be taken from the database. Default limit will acquire all results from the database")
	JSONCmd.Flags().IntVarP(&offset, "offset", "o", offset, "Offset is the number of records that get skipped during export.")
	JSONCmd.Flags().StringVarP(&filePath, "filepath", "f", filePath, "FilePath -f myfilepath")
	JSONCmd.MarkFlagsRequiredTogether("raw-sql", "raw-sql-filepath")

	cli := configs.NewCLIConfig()
	JSONCmd.Flags().VisitAll(cli.ConfigureViper("exporter"))
}

func ExportJSON(cmd *cobra.Command, args []string) error {
	var data interface{}
	var err error

	db := models.GetDB()
	exporterConfig := configs.LoadExporterConfig()
	setExportOptions(&exporterConfig, cmd)

	if !rawSQL {
		color.HiMagenta("[+] Constructing query...")
		qb := querybuilder.NewQueryBuilder(
			exporterConfig.Criteria,
			exporterConfig.Associations,
			exporterConfig.SortBy,
			exporterConfig.SortOrder,
			exporterConfig.Limit,
			exporterConfig.Offset,
		)
		data = qb.ConstructQuery(db)
	} else {
		color.HiRed("[+] Reading raw query...")
		qb := querybuilder.NewQueryBuilder(nil, "", "", "", 0, 0)
		err, data = qb.Raw(db, rawFilePath)
		if err != nil {
			return err
		}
	}
	err = exporters.ExportDataToJSON(data, exporterConfig.Filepath)
	if err != nil {
		return err
	}
	return nil
}

func setExportOptions(eC *configs.Exporter, cmd *cobra.Command) {
	if cmd.Flags().Changed("criteria") && len(criteria) != 0 {
		eC.Criteria = make(map[string]interface{})
		for key, value := range criteria {
			eC.Criteria[key] = value
		}
	}
	if cmd.Flags().Changed("associations") && associations != "" {
		eC.Associations = associations
	}
	if cmd.Flags().Changed("sort-by") && sortBy != "" {
		eC.SortBy = sortBy
	}

	if cmd.Flags().Changed("sort-order") && sortOrder != "" {
		eC.SortOrder = sortOrder
	}

	if cmd.Flags().Changed("limit") && limit >= 0 {
		eC.Limit = limit
	}

	if cmd.Flags().Changed("filepath") && filePath != "" {
		eC.Filepath = filePath
	}

}
