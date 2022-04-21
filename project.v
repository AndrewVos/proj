module main

import time
import os

struct Project {
mut:
	path     string
	name     string
	date     time.Time
	complete bool
	contents string
}

fn load_project(path string) Project {
	front_matter := retrieve_front_matter(path) or { panic(err) }
	contents := retrieve_contents(path)

	mut project := Project{
		path: path
		contents: contents
	}

	for line in front_matter.split('\n') {
		key := line.all_before('=')
		value := line.all_after('=')

		if key == 'name' {
			project.name = value
		} else if key == 'date' {
			project.date = time.parse(value) or { panic(err) }
		} else if key == 'complete' {
			project.complete = string_to_bool(value)
		}
	}

	return project
}

fn build_new_project(name string) Project {
	path := new_project_path(name) or {
		panic("Couldn't generate a project file path because all options already exist")
	}

	return Project{
		path: path
		name: name
		date: time.now()
		complete: false
		contents: ['# $name', '', '## Description', '', '## Tasks', '', '- [ ] First Task'].join('\n')
	}
}

fn (project Project) save() {
	date := project.date.format_ss()

	safe_write_file(project.path, [
		'---',
		'name=$project.name',
		'date=$date',
		'complete=$project.complete',
		'---',
		project.contents,
	].join('\n'))
}

fn (project Project) open_in_editor() {
	os.system([editor(), project.path].join(' '))
}

fn safe_write_file(path string, contents string) {
	tmp := path + '.tmp'

	mut file := os.create(tmp) or { panic(err) }
	file.write_string(contents) or { panic(err) }
	file.close()

	if os.is_file(path) {
		os.rm(path) or { panic(err) }
	}
	os.mv(tmp, path) or { panic(err) }
}

fn retrieve_contents(path string) string {
	file_contents := os.read_file(path) or { panic(err) }

	front_matter_start := file_contents.index('---') or {
		panic("can't find any front matter for file $path")
	}
	front_matter_end := file_contents.index_after('---', front_matter_start + 1)

	return file_contents[front_matter_end + 3..].trim_space()
}

fn retrieve_front_matter(path string) ?string {
	lines := os.read_lines(path) or { panic(err) }
	mut front_matter := []string{}

	mut started := false

	for line in lines {
		is_border := line.trim_space() == '---'

		if is_border {
			if !started {
				started = true
			} else {
				break
			}
		}

		front_matter << line
	}

	return front_matter.join('\n')
}

fn find_project_path_by_id(id string) ?string {
	project_paths := list_incomplete_project_paths()

	for index, project_path in project_paths {
		number := index + 1
		if number.str() == id {
			return project_path
		}
	}

	println("can't find project with id \"$id\"")
	exit(1)
}

fn new_project_path(name string) ?string {
	file_type := 'md'

	slug := generate_slug(name)

	for number := 0; number < 100; number++ {
		mut file_name := [slug, '.', file_type].join('')
		if number > 0 {
			file_name = [slug, '-', number.str(), '.', file_type].join('')
		}

		file_path := os.join_path(os.home_dir(), '.projects', file_name)

		if !os.is_file(file_path) {
			return file_path
		}
	}

	return none
}

fn list_incomplete_project_paths() []string {
	return list_project_paths().filter(fn (path string) bool {
		project := load_project(path)
		return !project.complete
	})
}

fn list_project_paths() []string {
	file_type := 'md'

	return os.glob(os.join_path(os.home_dir(), '.projects', '*.' + file_type)) or { return [] }
}
