package cmd

import (
	"github.com/AndrewVos/proj/project"
	"github.com/spf13/cobra"
	"log"
)

var listAllCmd = &cobra.Command{
	Use:   "list-all",
	Short: "List all projects",
	Run: func(cmd *cobra.Command, args []string) {
		projects, err := project.ListProjects()
		if err != nil {
			log.Fatal(err)
		}

		printProjects(projects, true)
	},
}

func init() {
	rootCmd.AddCommand(listAllCmd)
}
