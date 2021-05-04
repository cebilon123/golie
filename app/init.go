package app

import (
	commands "github.com/cebilon123/golie/app/cmd"
	"github.com/spf13/cobra"
)

func Init() {
	rootCmd := &cobra.Command{
		Use: "golie",
		Long: "Golie is distributed file share/sync CLI, which makes your memory lightweight and let you to store files" +
			"safely on different devices as long as they are paired with each other.",
	}

	rootCmd.AddCommand(
		config(),
	)

	if err := rootCmd.Execute(); err != nil {
		println(err)
	}
}

var config = func() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "config [-p path]",
		Long: "Manage golie configuration",
		Short: "Manage golie configuration",
		Run: func(cmd *cobra.Command, args []string) {
			if flag := cmd.Flag("path"); flag != nil {
				commands.SetPath(flag.Value.String())
			}
		},
	}

	return cmd
}
