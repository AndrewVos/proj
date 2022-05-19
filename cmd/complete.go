package cmd

import (
	"github.com/AndrewVos/proj/project"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"time"
)

var completeCmd = &cobra.Command{
	Use:   "complete <id>",
	Short: "Mark a project as complete",
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

		now := time.Now()
		project.Complete = &now
		err = project.Save()
		if err != nil {
			log.Fatal(err)
		}
	},
	ValidArgsFunction: validArgsFunction,
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
