package tests

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/spf13/viper"
)

var (
	_, b, _, _ = runtime.Caller(0)
	rootPath   = filepath.Join(filepath.Dir(b), "../..")
)

type Database struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"pass"`
	TestName string `mapstructure:"testing_name"`
	URL      string
}

func ReadTestConfig(configName, path, key string, cfg interface{}) error {
	if configName != "" {
		viper.SetConfigName(configName)
	} else {
		viper.SetConfigName("pryingdeep")
	}
	viper.SetConfigType("yaml")
	if path != "" {
		viper.AddConfigPath(path)
	}

	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.pryingdeep")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.UnmarshalKey(key, cfg); err != nil {
		return err
	}

	return nil
}

func TestYamlConfig(t *testing.T) {
	testCases := []struct {
		name     string
		config   string
		expected interface{}
	}{
		{
			name:     "Default Configuration",
			config:   "",
			expected: Database{Host: "localhost", Port: "5432", Name: "deep", User: "admin", Password: "admin", TestName: "testing"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var db Database
			err := ReadTestConfig(tc.config, rootPath, "database", &db)
			if err != nil {
				t.Fatalf("Failed to read test config %s", err)
			}

			if db != tc.expected {
				t.Errorf("Expected: %v, Got: %v", tc.expected, db)
			}
		})
	}
}
