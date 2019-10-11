package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/svenfinke/edm/manager"
)

func init(){
	rootCmd.AddCommand(fetchCmd)
}

var (
	fetchCmd = &cobra.Command{
		Use: "fetch",
		Short: "Fetch all defined dependencies",
		Long: "Fetch all defined dependencies",
		Run: func(cmd *cobra.Command, args []string) {
			var cfg = manager.OpenConfig(cfgFilename)

			for _, dep := range cfg.Dependencies {
				if err := dep.Fetch(); err != nil {
					fmt.Printf("ERR: Could not download or save file. (%s: %s)\n", dep.Target, err)
				}
			}
		},
	}
)