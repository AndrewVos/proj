module main

import os
import term
import strconv

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

	for project in projects {
		if all || !project.complete {
			mut complete_icon := term.colorize(term.green, '[x]')
			if !project.complete {
				complete_icon = term.colorize(term.red, '[ ]')
			}

			table << [term.colorize(term.blue, '#$project.id'), complete_icon, project.name,
				term.colorize(term.yellow, project.date.format_ss())]
		}
	}

	render_table(table)
}

fn edit(id int) {
	project := find_project(id) or { panic(err) }
	project.open_in_editor()
}

fn complete(id int) {
	mut project := find_project(id) or { panic(err) }
	project.complete = true
	project.save()
}

fn incomplete(id int) {
	mut project := find_project(id) or { panic(err) }
	project.complete = false
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
			id := strconv.atoi(args[1]) or { panic(err) }
			edit(id)
			return
		} else if args[0] == 'list' {
			list(false)
			return
		} else if args[0] == 'list-all' {
			list(true)
			return
		} else if args[0] == 'complete' && args.len == 2 {
			id := strconv.atoi(args[1]) or { panic(err) }
			complete(id)
			return
		} else if args[0] == 'incomplete' && args.len == 2 {
			id := strconv.atoi(args[1]) or { panic(err) }
			incomplete(id)
			return
		}
	}

	help()
}
