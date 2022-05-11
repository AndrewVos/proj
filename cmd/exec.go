package cmd

import (
	"github.com/AndrewVos/proj/project"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

var execCmd = &cobra.Command{
	Use:   "exec <id>",
	Short: "Execute code blocks inside your project files",
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

		err = project.Execute()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(execCmd)
}
