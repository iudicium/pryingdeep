package configs

type ExporterConfig struct {
	//Criteria is the map needed for exporting data.
	//Leave blank to not apply any criteria.
	Criteria map[string]interface{}
	//Associations are database tables that you can specify during export.
	//E.G all - default, will take all the tables.
	Associations string
	//SortBy is the ORDER BY field in web_pages
	SortBy string
	//SortOrder is the ASC,DESC value
	SortOrder string
	//Limit is the number of rows that you want to export
	Limit int
	//Filepath is where you would like to save your .json exporter data.
	//Default is data.json in the pryingdeep directory
	Filepath string
}

func LoadExporterConfig(path string) ExporterConfig {
	var config ExporterConfig
	loadConfig(path, &config)
	return config
}
