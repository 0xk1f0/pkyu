package cmd

import (
	"fmt"
	"pkyu/internal"

	"github.com/spf13/cobra"
)

var showVersion bool
var rootCmd = &cobra.Command{
	Use:   "pkyu",
	Short: "pkyu - Podman Kubernetes YAML Utility",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	Version: "0.1.0",
	Run: func(cmd *cobra.Command, args []string) {
		if showVersion {
		    fmt.Println("pkyu 0.1.0")
		} else {
		    cmd.Help()
		}
	},
}

func init() {
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
	    internal.ExitError(err, 1)
	}
}
