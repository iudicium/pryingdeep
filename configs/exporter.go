package configs

type Exporter struct {
	// Criteria is the map needed for exporting data.
	// Leave blank to not apply any criteria.
	Criteria map[string]interface{} `mapstructure:"criteria"`
	// Associations are database tables that you can specify during export.
	// E.G all - default, will take all the tables.
	Associations string `mapstructure:"associations"`
	// SortBy is the ORDER BY field in web_pages
	SortBy string `mapstructure:"sort-by"`
	// SortOrder is the ASC, DESC value
	SortOrder string `mapstructure:"sort-order"`
	// Limit is the number of rows that you want to export
	Limit int `mapstructure:"limit"`

	Offset int `mapstructure:"offset"`
	// Filepath is where you would like to save your .json exporter data.
	// Default is data.json in the current directory.
	Filepath string `mapstructure:"filepath"`
}

func LoadExporterConfig() Exporter {
	var config Exporter
	loadConfig("exporter", &config)
	return config
}
