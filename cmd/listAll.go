package cmd

import (
	"github.com/AndrewVos/proj/project"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
)

var listAllCmd = &cobra.Command{
	Use:   "list-all",
	Short: "List all projects",
	Run: func(cmd *cobra.Command, args []string) {
		projects, err := project.ListProjects()
		if err != nil {
			log.Fatal(err)
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetBorder(false)
		table.SetAutoWrapText(false)
		table.SetCenterSeparator("")
		table.SetColumnSeparator("")

		for _, project := range projects {
			completeIcon := "[x]"
			if !project.Complete {
				completeIcon = "[ ]"
			}

			table.Append(
				[]string{"#" + strconv.Itoa(project.ID), completeIcon, project.Name,
					project.Date.Format("2006-01-02 15:04:05"),
				},
			)
		}

		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(listAllCmd)
}
