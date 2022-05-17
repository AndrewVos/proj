package cmd

import (
	"fmt"
	"github.com/AndrewVos/proj/project"
	"github.com/justincampbell/timeago"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"golang.org/x/sys/unix"
	"log"
	"os"
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

func isTTY() bool {
	_, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ)
	return err == nil
}

func printProjects(projects []project.Project) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorder(false)
	table.SetAutoWrapText(false)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")

	for _, project := range projects {
		if !project.Complete || All {
			completeIcon := "[ ]"
			completeColour := tablewriter.Colors{tablewriter.Normal, tablewriter.FgRedColor}

			if project.Complete {
				completeIcon = "[x]"
				completeColour = tablewriter.Colors{tablewriter.Normal, tablewriter.FgGreenColor}
			}

			completionStatus := ""
			checklistCompletionColour := tablewriter.Colors{tablewriter.Normal, tablewriter.FgRedColor}
			if project.TasksTotal != 0 {
				completionStatus = fmt.Sprintf("%v/%v", project.TasksComplete, project.TasksTotal)
				if project.TasksComplete == project.TasksTotal {
					checklistCompletionColour = tablewriter.Colors{tablewriter.Normal, tablewriter.FgGreenColor}
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

			if isTTY() {
				table.Rich(
					cells,
					[]tablewriter.Colors{
						tablewriter.Colors{tablewriter.Normal, tablewriter.FgBlueColor},
						checklistCompletionColour,
						completeColour,
						tablewriter.Colors{},
						tablewriter.Colors{tablewriter.Normal, tablewriter.FgYellowColor},
					},
				)
			} else {
				table.Append(cells)
			}
		}
	}

	table.Render()
}
