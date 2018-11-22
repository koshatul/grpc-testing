package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use: "grpc-testing",
}

func init() {
	cobra.OnInitialize(configInit)

	rootCmd.PersistentFlags().BoolP("debug", "d", false, "Debug output")
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
	viper.BindEnv("debug", "DEBUG")
}

func main() {
	rootCmd.Execute()
}
