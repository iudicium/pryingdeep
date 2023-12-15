package install

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/iudicium/pryingdeep/pkg/fsutils"
	"github.com/iudicium/pryingdeep/pkg/logger"
)

const (
	configURL      = "https://raw.githubusercontent.com/iudicium/pryingdeep/main/pryingdeep.yaml"
	defaultCfgPath = ".pryingdeep/pryingdeep.yaml"
)

var InstallCMD = &cobra.Command{

	Use:   "install",
	Short: "Installation of config files",
	//This is needed to override the root.go PersitentPreRunE and not initialize the config,
	//since it doesn't exist yet
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	RunE: install,
}

func install(cmd *cobra.Command, args []string) error {
	fmt.Println("Welcome to prying deep! Let's begin the installation!")

	homefolder, err := os.UserHomeDir()
	if err != nil {
		logger.Errorf("Error getting home directory: %s", err)
		return err
	}

	cfgPath := path.Join(homefolder, defaultCfgPath)

	if fsutils.Exists(cfgPath) {
		color.Red("Config file already exists. Exiting installation.")
		return nil
	}

	fmt.Println("Config will be saved to:", cfgPath)
	if err := downloadConfigAndWriteConfig(cfgPath); err != nil {
		logger.Errorf("Error during installation: %s", err)
		return err
	}

	fmt.Println(color.GreenString("Config downloaded successfully!"))
	ctx := context.WithValue(context.Background(), "init", true)
	cmd.SetContext(ctx)
	return nil
}

func downloadConfigAndWriteConfig(filePath string) error {
	_ = fsutils.Touch(filePath)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(configURL)
	if err != nil {
		return fmt.Errorf("Error making HTTP request: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP request failed with status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Errorf("Error reading response body: %s", err)
		return err
	}

	err = fsutils.WriteTextFile(filePath, string(body))
	if err != nil {
		logger.Errorf("Error writing config file: %s", err)
		return err
	}
	return nil
}
