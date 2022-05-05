package cmd

import (
	"github.com/AndrewVos/proj/project"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

var completeCmd = &cobra.Command{
	Use:   "complete <id>",
	Short: "Mark a project complete",
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

		project.Complete = true
		err = project.Save()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
