package cmd

import (
	"os"
	"path/filepath"
	"pkyu/internal"

	"github.com/spf13/cobra"
)

var (
	doPruneAfter bool
)

func init() {
	downCmd.Flags().BoolVar(&doPruneAfter, "prune", false, "Prune unused data after removal")
	rootCmd.AddCommand(downCmd)
}

var downCmd = &cobra.Command{
	Use:   "down <FILE>",
	Short: "Remove pods based on Kubernetes YAML",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var (
			err          error
			kubeFilePath string
		)

		if kubeFilePath, err = filepath.Abs(args[0]); err != nil {
			internal.ExitError(err, 1)
		}

		if _, err = internal.ReadKubefile(kubeFilePath); err != nil {
			internal.ExitError(err, 1)
		}

		if !internal.BinaryExists("podman") {
			internal.ExitError("podman unavailable in $PATH", 1)
		}

		if msg, exitCode, err := internal.RunCommand("podman", "kube", "down", kubeFilePath).Single(); err != nil {
			internal.ExitError(msg, exitCode)
		}

		if doPruneAfter {
			if msg, exitCode, err := internal.RunCommand("podman", "system", "prune", "-f").Single(); err != nil {
				internal.ExitError(msg, exitCode)
			}
		}

		os.Exit(0)
	},
}
