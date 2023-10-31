package json

import (
	"log"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/pryingbytez/prying-deep/configs"
	"github.com/pryingbytez/prying-deep/models"
	"github.com/pryingbytez/prying-deep/pkg/exporters"
	"github.com/pryingbytez/prying-deep/pkg/logger"
	"github.com/pryingbytez/prying-deep/pkg/querybuilder"
)

var JSONCmd = &cobra.Command{
	Use:   "json",
	Short: "Export the crawled data to json",
	PreRun: func(cmd *cobra.Command, args []string) {
		if rawSQL {
			cmd.MarkFlagRequired(rawFilePath)
		}
	},
	Run: ExportJSON,
}

var (
	silent       = false
	rawSQL       = false
	rawFilePath  = "pkg/querybuilder/queries/select.sql"
	criteria     map[string]string
	config       = "configs/json/exporterConfig.json"
	associations = "all"
	sortBy       = "status_code"
	sortOrder    = "asc"
	limit        = 0
	filePath     = "data.json"
)

func init() {

	JSONCmd.PersistentFlags().BoolVarP(&silent, "silent", "s", silent, "-s to disable logging and run silently")
	JSONCmd.Flags().StringVarP(&config, "config", "c", config, "Config filepath -c myfilepath")
	JSONCmd.Flags().BoolVarP(&rawSQL, "rawSQL", "r", rawSQL, "--raw to use raw sql queries that you provide. All other flags except silent, rawFilePath and filepath will not matter.")
	JSONCmd.Flags().StringVarP(&rawFilePath, "rawSQLFilePath", "p", rawFilePath, "-rp to specify the file path to the sql file. Only use this flag if you specify -raw")
	JSONCmd.Flags().StringToStringVarP(&criteria, "criteria", "q", criteria, "JSON-formatted criteria, e.g., -c 'title=test,\"url=LIKE example.com\"'")
	JSONCmd.Flags().StringVarP(&associations, "associations", "a", associations, "-a WP,E,P,C")
	JSONCmd.Flags().StringVarP(&sortBy, "sortBy", "b", sortBy, "SortBy e.g -> -b title")
	JSONCmd.Flags().StringVarP(&sortOrder, "sortOrder", "o", sortOrder, "SortOrder e.g -> -o ASC || -b DESC. Only use this flag if you use SortBy")
	JSONCmd.Flags().IntVarP(&limit, "limit", "l", limit, "Limit e.g -> -l 100 -> 100 items will be taken from the database. Default limit will acquire all results from the database")
	JSONCmd.Flags().StringVarP(&filePath, "filePath", "f", filePath, "FilePath -f myfilepath")

}

func ExportJSON(cmd *cobra.Command, args []string) {
	var data interface{}
	var err error
	configs.LoadEnv()
	configs.LoadDatabase()
	logger.InitLogger(silent)
	defer logger.Logger.Sync()

	cfg := configs.GetConfig()
	db := models.SetupDatabase(cfg.DbConf.DbURL)

	if !rawSQL {
		viper.Set("Associations", associations)
		viper.Set("SortBy", sortBy)
		viper.Set("SortOrder", sortOrder)
		viper.Set("Limit", limit)
		viper.Set("FilePath", filePath)
		viper.Set("Criteria", criteria)

		configs.SaveConfig(config)

	}

	exporterConfig := configs.LoadExporterConfig(config)

	if !rawSQL {
		color.HiMagenta("[+] Constructing query...")
		qb := querybuilder.NewQueryBuilder(
			exporterConfig.Criteria,
			exporterConfig.Associations,
			exporterConfig.SortBy,
			exporterConfig.SortOrder,
			exporterConfig.Limit,
		)
		data = qb.ConstructQuery(db)
	} else {
		color.HiRed("[+] Reading raw query...")
		qb := querybuilder.NewQueryBuilder(nil, "", "", "", 0)
		err, data = qb.Raw(db, rawFilePath)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = exporters.ExportDataToJSON(data, exporterConfig.Filepath)
	if err != nil {
		log.Fatal(err)
	}
}
