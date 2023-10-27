package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pryingdeep",
	Short: "Pryingdeep is a dark web osint intelligence tool.",
	Long: `Pryingdeep specializes in collecting information about dark-web/clearnet websites.
		This tool was specifically built to extract as much information as possible from a .onion website`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// TODO: change the art i dont like it
func art() string {
	return `
$$$$$$$\  $$$$$$$\$$\     $$\$$$$$$\ $$\   $$\  $$$$$$\  $$$$$$$\  $$$$$$$$\ $$$$$$$$\ $$$$$$$\
$$  __$$\ $$  __$$\$$\   $$  \_$$  _|$$$\  $$ |$$  __$$\ $$  __$$\ $$  _____|$$  _____|$$  __$$\
$$ |  $$ |$$ |  $$ \$$\ $$  /  $$ |  $$$$\ $$ |$$ /  \__|$$ |  $$ |$$ |      $$ |      $$ |  $$ |
$$$$$$$  |$$$$$$$  |\$$$$  /   $$ |  $$ $$\$$ |$$ |$$$$\ $$ |  $$ |$$$$$\    $$$$$\    $$$$$$$  |
$$  ____/ $$  __$$<  \$$  /    $$ |  $$ \$$$$ |$$ |\_$$ |$$ |  $$ |$$  __|   $$  __|   $$  ____/
$$ |      $$ |  $$ |  $$ |     $$ |  $$ |\$$$ |$$ |  $$ |$$ |  $$ |$$ |      $$ |      $$ |
$$ |      $$ |  $$ |  $$ |   $$$$$$\ $$ | \$$ |\$$$$$$  |$$$$$$$  |$$$$$$$$\ $$$$$$$$\ $$ |
\__|      \__|  \__|  \__|   \______|\__|  \__| \______/ \_______/ \________|\________|\__|
`
}
