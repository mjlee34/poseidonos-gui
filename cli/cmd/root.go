package cmd

import (
		"os"
		"github.com/spf13/cobra"
	   )

var Verbose bool
var Debug bool

var IP string
var Port string

var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "IBoF cli is simple client of IBoF.",
	Long: `IBoF cli is simple client of IBoF.

You can set ip and port number of iBoF using config.yaml and use a "command" command with -i flag.
		`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().BoolVarP(&Debug, "debug", "d", false, "set a debug mode")
	
	rootCmd.PersistentFlags().StringVar(&IP, "ip", "", "set ip adddress like \"--ip 127.0.0.1\"")
	rootCmd.PersistentFlags().StringVar(&Port, "port", "", "set port number like \"--port 18716\"")
}