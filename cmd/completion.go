package cmd

import (
	"fmt"
	"github.com/AndrewVos/proj/project"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
)

func validArgsFunction(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if len(args) != 0 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	projectIds, err := projectIds(toComplete)
	if err != nil {
		log.Fatal(err)
	}
	return projectIds, cobra.ShellCompDirectiveNoFileComp
}

func projectIds(filter string) ([]string, error) {
	projects, err := project.ListProjects()
	if err != nil {
		return nil, err
	}

	result := []string{}

	for _, project := range projects {
		if project.Visible() && projectMatchesFilter(project, filter) {
			result = append(result, fmt.Sprintf("%v\t%v", project.ID, project.Name))
		}
	}

	return result, nil
}

func projectMatchesFilter(project project.Project, filter string) bool {
	id := strconv.Itoa(project.ID)

	if strings.HasPrefix(id, filter) {
		return true
	}

	if strings.Contains(strings.ToLower(project.Name), strings.ToLower(filter)) {
		return true
	}

	return false
}
