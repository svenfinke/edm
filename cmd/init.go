package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/svenfinke/edm/manager"
)

func init() {
	initCmd.Flags().BoolVarP(&withExample, "withExample", "e", false, "adds examples to the initial file")
	rootCmd.AddCommand(initCmd)
}

var (
	withExample bool

	initCmd = &cobra.Command{
		Use:   "init",
		Short: "Initialize the config file for edm",
		Long: `Init will create a .edm.yml file in the root of the project. You can alter the location or name of the config file by passing flags to the function. 

Example usage:
	edm init

`,
		Run: func(cmd *cobra.Command, args []string) {
			var cfg = manager.Config{}

			if withExample {
				var depExample1 = manager.Dependency{
					Target: "./config/some.yaml",
					Source: "https://example.com/some.yaml",
					Type:   "default",
				}
				var depExample2 = manager.Dependency{
					Target: "./.env",
					Source: "https://example.com/.env",
					Type:   "file",
				}

				cfg.Dependencies = append(cfg.Dependencies, &depExample1)
				cfg.Dependencies = append(cfg.Dependencies, &depExample2)
			}

			if err := cfg.WriteFile(cfgFilename); err != nil {
				fmt.Printf("\nERR: error while creating file. (%s)", err.Error())
			}
		},
	}
)
