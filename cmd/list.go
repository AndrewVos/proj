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

var All bool
var Date bool

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
	listCmd.Flags().BoolVarP(&Date, "date", "d", false, "show full date")

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

func (c ChecklistCompletionCell) percentage() int {
	if c.Project.TasksComplete == 0 {
		return 0
	}
	return c.Project.TasksComplete / c.Project.TasksTotal * 100
}

func (c ChecklistCompletionCell) Width() int {
	return len(strconv.Itoa(c.percentage())) + 1
}

func (c ChecklistCompletionCell) Render() {
	if c.Project.TasksComplete == c.Project.TasksTotal {
		color.New(color.FgGreen).Printf("%v", c.percentage())
		color.New(color.FgGreen).Print("%")
	} else {
		color.New(color.FgRed).Printf("%v", c.percentage())
		color.New(color.FgRed).Print("%")
	}
}

type DateStatusCell struct {
	Project project.Project
}

func (c DateStatusCell) formatDate() string {
	if Date {
		return c.Project.Date.Format("2006-01-02 15:04")
	}
	return timeago.FromDuration(time.Since(c.Project.Date)) + " old"
}

func (c DateStatusCell) Width() int {
	return len(c.formatDate())
}

func (c DateStatusCell) Render() {
	color.New(color.FgMagenta).Printf(c.formatDate())
}

type TitleCell struct {
	Project project.Project
}

func (c TitleCell) Width() int {
	return len(c.Project.Name)
}

func (c TitleCell) Render() {
	if c.Project.Complete {
		color.New(color.CrossedOut).Print(c.Project.Name)
	} else {
		fmt.Printf(c.Project.Name)
	}
}

func printProjects(projects []project.Project) {
	t := table.New()
	t.SetCellStretch(2)
	t.SetCellAlignment(3, table.AlignRight)

	for _, project := range projects {
		if !project.Complete || All {
			cells := []table.Cell{
				IdCell{project},
				ChecklistCompletionCell{project},
				TitleCell{project},
				DateStatusCell{project},
			}

			t.Row(cells)
		}
	}

	t.Print()
}
