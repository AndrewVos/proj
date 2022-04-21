module main

import os

fn help() {
	println('proj')
	println('  Usage: proj create "Project Name"')
	println('  Usage: proj list')
	println('  Usage: proj list-all')
	println('  Usage: proj edit <id>')
	println('  Usage: proj complete <id>')
}

fn create(name string) {
	project := build_new_project(name)
	project.save()
	project.open_in_editor()
}

fn list() {
	projects := list_projects()

	mut table := [['ID', 'Name', 'Date']]

	for index, project in projects {
		if !project.complete {
			number := index + 1
			table << [number.str(), project.name, project.date.format_ss()]
		}
	}

	render_table(table)
}

fn list_all() {
	projects := list_projects()

	mut table := [['ID', 'Name', 'Date']]

	for index, project in projects {
		number := index + 1
		table << [number.str(), project.name, project.date.format_ss()]
	}

	render_table(table)
}

fn edit(id string) {
	project := find_project(id) or { panic(err) }
	project.open_in_editor()
}

fn complete(id string) {
	mut project := find_project(id) or { panic(err) }
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
		} else if args[0] == 'list-all' {
			list_all()
			return
		} else if args[0] == 'complete' && args.len == 2 {
			complete(args[1])
			return
		}
	}

	help()
}
