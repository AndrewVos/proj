package cmd

import (
	"github.com/AndrewVos/proj/project"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"time"
)

var cancelCmd = &cobra.Command{
	Use:   "cancel <id>",
	Short: "Mark a project as cancelled",
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
		project.Cancelled = &now
		err = project.Save()
		if err != nil {
			log.Fatal(err)
		}
	},
	ValidArgsFunction: validArgsFunction,
}

func init() {
	rootCmd.AddCommand(cancelCmd)
}
