/*
Copyright © 2025 Sina Sadeghi sina.sadeghi83@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/Ahu-Tools/AhuM/cmd/edge"
	"github.com/Ahu-Tools/AhuM/cmd/gin"
	"github.com/Ahu-Tools/AhuM/cmd/initialise"
	"github.com/Ahu-Tools/AhuM/cmd/service"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ahum",
	Short: "A project management wizard for Ahu services",
	Long:  `This wizard will help you to create and manage an Ahu project. You can choose modules you need for your project and your porject will be shipped with the modules you need.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(initialise.InitCmd)
	rootCmd.AddCommand(edge.EdgeCmd)
	rootCmd.AddCommand(service.ServiceCmd)
	rootCmd.AddCommand(gin.GinCmd)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ahum.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".ahum" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".ahum")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
