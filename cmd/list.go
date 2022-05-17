package cmd

import (
	"fmt"
	"github.com/AndrewVos/proj/project"
	"github.com/AndrewVos/proj/table"
	"github.com/fatih/color"
	"github.com/justincampbell/timeago"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"time"
)

var Relative bool
var All bool

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List projects",
	Run: func(cmd *cobra.Command, args []string) {
		projects, err := project.ListProjects()
		if err != nil {
			log.Fatal(err)
		}

		printProjects(projects)
	},
}

func init() {
	listCmd.Flags().BoolVarP(&All, "all", "a", false, "list all projects")
	listCmd.Flags().BoolVarP(&Relative, "relative", "r", false, "relative time output")

	rootCmd.AddCommand(listCmd)
}

func printProjects(projects []project.Project) {
	table := table.New()

	for _, project := range projects {
		if !project.Complete || All {
			completeIcon := "[ ]"
			completeColour := color.New(color.FgRed)

			if project.Complete {
				completeIcon = "[x]"
				completeColour = color.New(color.FgGreen)
			}

			completionStatus := ""
			checklistCompletionColour := color.New(color.FgRed)
			if project.TasksTotal != 0 {
				completionStatus = fmt.Sprintf("%v/%v", project.TasksComplete, project.TasksTotal)
				if project.TasksComplete == project.TasksTotal {
					checklistCompletionColour = color.New(color.FgGreen)
				}
			}

			formattedDate := project.Date.Format("2006-01-02 15:04")

			if Relative {
				formattedDate = timeago.FromDuration(time.Since(project.Date)) + " ago"
			}

			cells := []string{
				"#" + strconv.Itoa(project.ID),
				completionStatus,
				completeIcon,
				project.Name,
				formattedDate,
			}

			table.Row(
				cells,
			)

			table.ColouriseRow([]*color.Color{color.New(color.FgBlue),
				checklistCompletionColour,
				completeColour,
				color.New(),
				color.New(color.FgYellow),
			})
		}
	}

	table.Print()
}
