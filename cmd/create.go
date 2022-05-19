package cmd

import (
	"github.com/AndrewVos/proj/project"
	"github.com/spf13/cobra"
	"log"
)

var Edit bool

var createCmd = &cobra.Command{
	Use:   "create <name>",
	Short: "Create a new project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		project, err := project.NewProject(args[0])
		if err != nil {
			log.Fatal(err)
		}
		err = project.Save()
		if err != nil {
			log.Fatal(err)
		}

		if Edit {
			err = project.OpenInEditor()
			if err != nil {
				log.Fatal(err)
			}
		}
	},
}

func init() {
	createCmd.Flags().BoolVarP(&Edit, "edit", "e", false, "edit the project immediately after creation")

	rootCmd.AddCommand(createCmd)
}
