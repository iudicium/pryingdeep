package configs

type ExporterConfig struct {
	Criteria     map[string]interface{}
	Associations string
	SortBy       string
	SortOrder    string
	Limit        int
	Filepath     string
}

func LoadExporterConfig(path string) ExporterConfig {
	var config ExporterConfig
	loadConfig(path, &config)
	return config
}
