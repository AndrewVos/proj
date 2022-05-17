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

type IdCell struct {
	Project project.Project
}

func (c IdCell) Width() int {
	return len("#" + strconv.Itoa(c.Project.ID))
}

func (c IdCell) Render() {
	color.New(color.FgBlue).Printf("#" + strconv.Itoa(c.Project.ID))
}

type ChecklistCompletionCell struct {
	Project project.Project
}

func (c ChecklistCompletionCell) Width() int {
	if c.Project.TasksTotal != 0 {
		result := fmt.Sprintf("%v/%v", c.Project.TasksComplete, c.Project.TasksTotal)
		return len(result)
	}
	return 0
}

func (c ChecklistCompletionCell) Render() {
	if c.Project.TasksTotal != 0 {
		result := fmt.Sprintf("%v/%v", c.Project.TasksComplete, c.Project.TasksTotal)
		if c.Project.TasksComplete == c.Project.TasksTotal {
			color.New(color.FgGreen).Printf(result)
		} else {
			color.New(color.FgRed).Printf(result)
		}
	}
}

type CompleteStatusCell struct {
	Project project.Project
}

func (c CompleteStatusCell) Width() int {
	return 3
}

func (c CompleteStatusCell) Render() {
	if c.Project.Complete {
		color.New(color.FgGreen).Printf("[x]")
	} else {
		color.New(color.FgRed).Printf("[ ]")
	}
}

type DateStatusCell struct {
	Project project.Project
}

func (c DateStatusCell) formatDate() string {
	if Relative {
		return timeago.FromDuration(time.Since(c.Project.Date)) + " ago"
	}
	return c.Project.Date.Format("2006-01-02 15:04")
}

func (c DateStatusCell) Width() int {
	return len(c.formatDate())
}

func (c DateStatusCell) Render() {
	color.New(color.FgYellow).Printf(c.formatDate())
}

func printProjects(projects []project.Project) {
	t := table.New()

	for _, project := range projects {
		if !project.Complete || All {
			cells := []table.Cell{
				IdCell{project},
				ChecklistCompletionCell{project},
				CompleteStatusCell{project},
				table.SimpleCell{project.Name},
				DateStatusCell{project},
			}

			t.Row(cells)
		}
	}

	t.Print()
}
