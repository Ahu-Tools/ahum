/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise an Ahu project",
	Long:  `Create folders and go files related to the main architecture of the Ahu project`,
	Args:  cobra.MatchAll(cobra.RangeArgs(1, 3), checkPortRange),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func checkPortRange(cmd *cobra.Command, args []string) error {
	var portStr string

	switch len(args) {
	case 2:
		portStr = args[1]
	case 3:
		portStr = args[2]
	default:
		return nil
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return errors.New("port must be a positive integer number")
	}

	if port > 65535 || port <= 0 {
		return errors.New("provided port is not in the valid range")
	}

	return nil
}
