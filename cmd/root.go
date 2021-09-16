package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "alto",
	Short: "Alto is a speedy, light package manager",
	Long:  `A speedy open source package manager that's ready for the masses, written in Go. Check it out at altopkg.com or github.com/altopm/alto`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
