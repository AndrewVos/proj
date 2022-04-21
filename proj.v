module main

import os
import term

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

fn list(all bool) {
	projects := list_projects()
	mut table := [][]string{}

	for index, project in projects {
		if all || !project.complete {
			number := index + 1

			mut complete_icon := term.colorize(term.red, '[ ]')
			if !project.complete {
				complete_icon = term.colorize(term.green, '[x]')
			}

			table << [term.colorize(term.blue, '#$number'), complete_icon, project.name,
				term.colorize(term.yellow, project.date.format_ss())]
		}
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
			list(false)
			return
		} else if args[0] == 'list-all' {
			list(true)
			return
		} else if args[0] == 'complete' && args.len == 2 {
			complete(args[1])
			return
		}
	}

	help()
}
