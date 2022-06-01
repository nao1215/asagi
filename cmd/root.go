package cmd

import (
	"github.com/nao1215/asagi/internal/completion"
	"github.com/nao1215/asagi/internal/print"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "asagi",
}

// Execute start command.
func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	completion.DeployShellCompletionFileIfNeeded(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		print.Fatal(err)
	}
}
