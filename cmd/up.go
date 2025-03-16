package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"pkyu/internal"

	"github.com/spf13/cobra"
)

var (
	needsBuild bool
	replaceOld bool
	pullNew    bool
)

func init() {
	upCmd.Flags().BoolVar(&replaceOld, "replace", false, "Delete then recreate pods in YAML")
	upCmd.Flags().BoolVar(&needsBuild, "build", false, "Build all images in a YAML")
	upCmd.Flags().BoolVar(&pullNew, "pull", false, "Build all images in a YAML")
	rootCmd.AddCommand(upCmd)
}

var upCmd = &cobra.Command{
	Use:   "up <FILE>",
	Short: "Play a pod or volume based on Kubernetes YAML",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var (
			err          error
			kubeFilePath string
			passCommands []string
		)

		kubeFilePath, err = filepath.Abs(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		_, err = internal.ReadKubefile(kubeFilePath)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if replaceOld {
			passCommands = append(passCommands, "--replace")
		}

		if !internal.BinaryExists("podman") {
			internal.ExitError("podman unavailable in $PATH", 1)
		}

		var fullExec = append([]string{"podman", "kube", "play"}, append(passCommands, kubeFilePath)...)
		if msg, exitCode, err := internal.RunCommand(fullExec...).Single(); err != nil {
			internal.ExitError(msg, exitCode)
		}

		os.Exit(0)
	},
}
