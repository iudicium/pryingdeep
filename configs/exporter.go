package configs

type ExporterConfig struct {
	WebPageCriteria map[string]interface{}
	Associations    string
	SortBy          string
	SortOrder       string
	Limit           int
	Filepath        string
}

func LoadExporterConfig() ExporterConfig {
	var config ExporterConfig
	loadConfig("exporterConfig.json", &config)
	return config

}
