package cmd

import (
	"github.com/AndrewVos/proj/project"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"golang.org/x/sys/unix"
	"log"
	"os"
	"strconv"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List projects",
	Run: func(cmd *cobra.Command, args []string) {
		projects, err := project.ListProjects()
		if err != nil {
			log.Fatal(err)
		}

		printProjects(projects, false)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func isTTY() bool {
	_, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ)
	return err == nil
}

func printProjects(projects []project.Project, all bool) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorder(false)
	table.SetAutoWrapText(false)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")

	for _, project := range projects {
		if !project.Complete || all {
			completeIcon := "[ ]"
			completeColour := tablewriter.Colors{tablewriter.Normal, tablewriter.FgRedColor}

			if project.Complete {
				completeIcon = "[x]"
				completeColour = tablewriter.Colors{tablewriter.Normal, tablewriter.FgGreenColor}
			}

			cells := []string{
				"#" + strconv.Itoa(project.ID),
				completeIcon,
				project.Name,
				project.Date.Format("2006-01-02 15:04:05"),
			}

			if isTTY() {
				table.Rich(
					cells,
					[]tablewriter.Colors{
						tablewriter.Colors{tablewriter.Normal, tablewriter.FgBlueColor},
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
