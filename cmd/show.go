package cmd

import (
	"github.com/AndrewVos/proj/project"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

var showCmd = &cobra.Command{
	Use:   "show <id>",
	Short: "Show the project contents",
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

		err = project.Show()
		if err != nil {
			log.Fatal(err)
		}
	},
	ValidArgsFunction: validArgsFunction,
}

func init() {
	rootCmd.AddCommand(showCmd)
}
