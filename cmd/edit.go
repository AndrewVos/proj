package cmd

import (
	"github.com/AndrewVos/proj/project"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

var editCmd = &cobra.Command{
	Use:   "edit <id>",
	Short: "Edit a project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}

		project, err := project.Find(id)
		if err != nil {
			log.Fatal(err)
		}

		err = project.OpenInEditor()
		if err != nil {
			log.Fatal(err)
		}
	},
	ValidArgsFunction: validArgsFunction,
}

func init() {
	rootCmd.AddCommand(editCmd)
}
