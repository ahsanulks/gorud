package cmd

import (
	"fmt"

	"github.com/ahsanulks/gorud/generate"
	"github.com/spf13/cobra"
)

const (
	defaultPackage string = "repository"
	defaultDir     string = "./repository/"
)

var (
	packageDestination string
	dirDestination     string
)

var cmdGenerate = &cobra.Command{
	Use:   "generate [file path]",
	Short: "Generate function create read update delete of struct",
	Long:  `Will be generate function create read update delete of struct`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return cmd.Help()
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return nil
		}
		g := generate.NewFilePath(args[0])
		err := g.ReadFile()
		if err != nil {
			return err
		}
		g.Write()
		fmt.Println("successfully generated")
		return nil
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	cmdGenerate.PersistentFlags().StringVarP(&packageDestination, "package", "p", "repository", "Package name.")
	cmdGenerate.PersistentFlags().StringVarP(&dirDestination, "directory", "d", "repository", "Directory destination.")
}

func initConfig() {
	packageDestination, _ := cmdGenerate.Flags().GetString("package")
	if packageDestination == "" {
		packageDestination = defaultPackage
	}
	dirDestination, _ := cmdGenerate.Flags().GetString("directory")
	if dirDestination == "" {
		dirDestination = defaultDir
	}
}
