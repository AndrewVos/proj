module main

import os
import term

fn help() {
	tool_name := term.colorize(term.green, 'proj')

	println(tool_name)
	println('')

	usage_lines := [
		['create', '"Project Name"'],
		['list'],
		['list-all'],
		['edit', '<id>'],
		['complete', '<id>'],
	]

	for line in usage_lines {
		print('  Usage: $tool_name ')
		for index, part in line {
			if index == 0 {
				print(term.colorize(term.blue, part))
			} else {
				print(' ')
				print(term.colorize(term.red, part))
			}
		}
		println('')
	}
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

			mut complete_icon := term.colorize(term.green, '[x]')
			if !project.complete {
				complete_icon = term.colorize(term.red, '[ ]')
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
