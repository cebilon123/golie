package app

import (
	"github.com/cebilon123/golie/app/cmd"
	"github.com/spf13/cobra"
)

func Init() {
	rootCmd := &cobra.Command{
		Use: "golie",
		Long: "Golie is distributed file share/sync CLI, which makes your memory lightweight and let you to store files" +
			"safely on different devices as long as they are paired with each other.",
	}

	rootCmd.AddCommand(cmd.GetConfigureCommand())

	if err := rootCmd.Execute(); err != nil {
		println(err)
	}
}
