package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/svenfinke/edm/manager"
)

func init() {
	rootCmd.AddCommand(fetchCmd)

	for name, dt := range manager.DependencyTypes {
		fetchCmd.Long += fmt.Sprintf("\n - %s: %s", name, dt.GetInfo())
	}
}

var (
	fetchCmd = &cobra.Command{
		Use:   "fetch",
		Short: "Fetch all defined dependencies",
		Long:  "Fetch all defined dependencies. Depending on the type of the dependency, the source or target might be used in different ways or additional actions might be taken. These are the available types that can be defined:",
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
