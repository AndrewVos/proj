module main

import os
import time

fn help() {
	println('proj')
	println('  Usage: proj create "Project Name"')
	println('  Usage: proj list')
	println('  Usage: proj edit <id>')
	println('  Usage: proj complete <id>')
}

fn create(name string) {
	path := new_project_path(name) or {
		panic("Couldn't generate a project file path because all options already exist")
	}

	date := time.now().format_ss()

	mut f := os.create(path) or { panic(err) }
	f.write_string(['---', 'name=$name', 'date=$date', '---', '', '# $name', '', '## Description',
		'', '## Tasks', '', '- [ ] First'].join('\n')) or { panic(err) }
	f.close()

	open_in_editor(path)
}

fn list() {
	project_paths := list_incomplete_project_paths()

	mut table := [['ID', 'Name', 'Date']]

	for index, project_path in project_paths {
		number := index + 1
		project := new_project(project_path)
		table << [number.str(), project.name, project.date.format_ss()]
	}

	render_table(table)
}

fn edit(id string) {
	path := find_project_path_by_id(id) or { panic(err) }
	open_in_editor(path)
}

fn complete(id string) {
	project_path := find_project_path_by_id(id) or { panic(err) }
	mut project := new_project(project_path)
	project.complete = true
	project.save()
}

fn main() {
	mut args := []string{}
	if os.args.len > 1 {
		args = os.args[1..]

		if args[0] == 'create' && args.len == 2 {
			create(args[1])
			return
		} else if args[0] == 'edit' && args.len == 2 {
			edit(args[1])
			return
		} else if args[0] == 'list' {
			list()
			return
		} else if args[0] == 'complete' && args.len == 2 {
			complete(args[1])
			return
		}
	}

	help()
}
