/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/AndrewVos/proj/project"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List projects",
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
			if !project.Complete {
				completeIcon := "[ ]"

				table.Append(
					[]string{"#" + strconv.Itoa(project.ID), completeIcon, project.Name,
						project.Date.Format("2006-01-02 15:04:05"),
					},
				)
			}
		}

		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
