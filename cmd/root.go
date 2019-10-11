package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func Execute(){
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init(){
	rootCmd.PersistentFlags().StringVarP(&cfgFilename, "filename", "f", ".edm.yaml", "Change the filename to use a different config file for edm.")
}

var (
	cfgFilename string

	rootCmd = &cobra.Command{
		Use:   "edm",
		Short: "External Dependency Manager",
		Long: `An easy to use manager for external dependencies of different kind in your project.`,
	}
)