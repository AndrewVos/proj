package project

import (
	"bufio"
	"errors"
	"github.com/AndrewVos/proj/markdown"
	"github.com/gosimple/slug"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Project struct {
	ID            int
	Path          string
	Name          string
	Date          time.Time
	Complete      *time.Time
	Contents      string
	TasksTotal    int
	TasksComplete int
}

func (p Project) Execute() error {
	snippets := markdown.FindSnippets(p.Contents)

	for _, snippet := range snippets {
		if snippet.Lang == "ruby --run" {
			cmd := exec.Command("ruby", "-e", snippet.Content)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			return cmd.Run()
		}
	}

	return nil
}

func NewProject(name string) (Project, error) {
	path, err := BuildProjectPath(name)

	if err != nil {
		return Project{}, err
	}

	id, err := NewProjectId()

	if err != nil {
		return Project{}, err
	}

	content := []string{
		"# " + name,
		"",
		"## Description",
		"",
		"## Tasks",
		"",
		"- [ ] First Task",
	}

	return Project{
		ID:       id,
		Path:     path,
		Name:     name,
		Date:     time.Now(),
		Contents: strings.Join(content, "\n"),
	}, nil
}

func BuildProjectPath(name string) (string, error) {
	fileType := "md"

	slug := slug.Make(name)

	for number := 0; number < 100; number++ {
		fileName := strings.Join([]string{slug, ".", fileType}, "")

		if number > 0 {
			fileName = strings.Join([]string{slug, "-", strconv.FormatInt(int64(number), 10), ".", fileType}, "")
		}

		homeDirectory, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}

		filePath := path.Join(homeDirectory, ".projects", fileName)

		if !FileExists(filePath) {
			return filePath, nil
		}
	}

	return "", errors.New("can't create filename because all options exist")
}

func NewProjectId() (int, error) {
	id := 0

	projects, err := ListProjects()
	if err != nil {
		return 0, err
	}
	for _, project := range projects {
		if project.ID > id {
			id = project.ID
		}
	}

	return id + 1, nil
}

func FileExists(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false
		}
		return false
	}
	return !fileInfo.IsDir()
}

func ListProjects() ([]Project, error) {
	fileType := "md"

	homeDirectory, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	projectPaths, err := filepath.Glob(path.Join(homeDirectory, ".projects", "*."+fileType))
	if err != nil {
		return nil, err
	}

	projects := []Project{}

	for _, projectPath := range projectPaths {
		project, err := LoadProject(projectPath)
		if err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}

	sort.Slice(projects, func(i, j int) bool {
		return projects[i].Date.Before(projects[j].Date)
	})

	return projects, nil
}

func LoadProject(path string) (Project, error) {
	frontMatter, err := retrieveFrontMatter(path)
	if err != nil {
		return Project{}, err
	}

	contents, err := retrieveContents(path)
	if err != nil {
		return Project{}, err
	}

	project := Project{
		Path:     path,
		Contents: contents,
	}

	for _, line := range strings.Split(frontMatter, "\n") {
		indexOfEqual := strings.Index(line, "=")
		if indexOfEqual < 0 {
			continue
		}

		key := line[0:indexOfEqual]
		value := line[indexOfEqual+1:]

		if key == "id" {
			id, err := strconv.Atoi(value)
			if err != nil {
				return Project{}, err
			}
			project.ID = id

		} else if key == "name" {
			project.Name = value
		} else if key == "date" {
			date, err := time.Parse(time.RFC3339, value)
			if err != nil {
				return Project{}, err
			}
			project.Date = date
		} else if key == "complete" {
			date, err := time.Parse(time.RFC3339, value)
			if err != nil {
				return Project{}, err
			}
			project.Complete = &date
		}
	}

	for _, line := range strings.Split(contents, "\n") {
		if strings.HasPrefix(strings.TrimSpace(line), "- [ ]") {
			project.TasksTotal += 1
		} else if strings.HasPrefix(strings.TrimSpace(line), "- [x]") {
			project.TasksTotal += 1
			project.TasksComplete += 1
		}
	}

	return project, nil
}

func stringToBool(s string) bool {
	return s == "true"
}

func retrieveFrontMatter(path string) (string, error) {
	lines, err := ReadLines(path)
	if err != nil {
		return "", err
	}
	frontMatter := []string{}

	started := false

	for _, line := range lines {
		isBorder := strings.TrimSpace(line) == "---"

		if isBorder {
			if !started {
				started = true
			} else {
				break
			}
		}

		frontMatter = append(frontMatter, line)
	}

	return strings.Join(frontMatter, "\n"), nil
}

func indexAt(s, sep string, n int) int {
	idx := strings.Index(s[n:], sep)
	if idx > -1 {
		idx += n
	}
	return idx
}

func retrieveContents(path string) (string, error) {
	contents, err := ReadFile(path)

	if err != nil {
		return "", err
	}

	frontMatterStart := strings.Index(contents, "---")
	frontMatterEnd := indexAt(contents, "---", frontMatterStart+1)

	return strings.TrimSpace(contents[frontMatterEnd+3:]), nil
}

func ReadFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	file.Close()

	return text, nil
}

func (p Project) Save() error {
	date := p.Date.Format(time.RFC3339)

	rows := []string{
		"---",
		"id=" + strconv.Itoa(p.ID),
		"name=" + p.Name,
		"date=" + date,
	}

	if p.Complete != nil {
		rows = append(rows, "complete="+p.Complete.Format(time.RFC3339))
	}
	rows = append(rows, "---")
	rows = append(rows, "")
	rows = append(rows, p.Contents)

	contents := strings.Join(rows, "\n") + "\n"

	return os.WriteFile(p.Path, []byte(contents), 0644)
}

func (p Project) OpenInEditor() error {
	editor := os.Getenv("EDITOR")

	cmd := exec.Command(editor, p.Path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func Find(id int) (Project, error) {
	projects, err := ListProjects()

	if err != nil {
		return Project{}, err
	}

	for _, project := range projects {
		if project.ID == id {
			return project, nil
		}
	}

	return Project{}, errors.New("can't find project with id \"" + strconv.Itoa(id) + "\"")
}

func (p Project) Show() error {
	cmd := exec.Command("glow", p.Path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
