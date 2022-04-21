module main

import os

fn help() {
	println('proj')
	println('  Usage: proj create "Project Name"')
	println('  Usage: proj list')
	println('  Usage: proj edit <id>')
	println('  Usage: proj complete <id>')
}

fn create(name string) {
	project := build_new_project(name)
	project.save()
	project.open_in_editor()
}

fn list() {
	project_paths := list_incomplete_project_paths()

	mut table := [['ID', 'Name', 'Date']]

	for index, project_path in project_paths {
		number := index + 1
		project := load_project(project_path)
		table << [number.str(), project.name, project.date.format_ss()]
	}

	render_table(table)
}

fn edit(id string) {
	path := find_project_path_by_id(id) or { panic(err) }
	project := load_project(path)
	project.open_in_editor()
}

fn complete(id string) {
	path := find_project_path_by_id(id) or { panic(err) }
	mut project := load_project(path)
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
