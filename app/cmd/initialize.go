package cmd

import "github.com/spf13/cobra"

func GetConfigureCommand() *cobra.Command {
	var (
		//Path to directory to watch
		watchDirectoryPath string
		//Device ip address with whom files are synced
		syncDeviceAddress string
	)

	command := &cobra.Command{
		Use:   "configure",
		Short: "Configure environment to use golie",
		Run: func(cmd *cobra.Command, args []string) {
			println(watchDirectoryPath)
		},
	}

	command.PersistentFlags().StringVar(&watchDirectoryPath, "path", "", "Path to directory which should be synced")
	command.PersistentFlags().StringVar(&syncDeviceAddress, "address", "192.168.100.100", "Address of device with whom files are synced")

	return command
}
