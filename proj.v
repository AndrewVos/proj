import os
import regex
import encoding.utf8
import time
import math

struct Project {
mut:
	name string
	date time.Time
}

fn help() {
	println('proj')
	println('  Usage: proj create "Project Name"')
	println('  Usage: proj list')
	println('  Usage: proj edit <id>')
}

fn list() {
	project_paths := list_project_paths()

	mut table := [['ID', 'Name', 'Date']]

	for index, project_path in project_paths {
		number := index + 1
		project := read_project(project_path)
		table << [number.str(), project.name, project.date.format_ss()]
	}

	render_table(table)
}

fn right_pad(s string, width int) string {
	mut new_string := s

	for {
		if new_string.len >= width {
			return new_string
		}
		new_string = new_string + ' '
	}

	return new_string
}

fn render_table(table [][]string) {
	mut column_widths := map[int]int{}

	for row in table {
		for column_index, cell in row {
			column_widths[column_index] = math.max(column_widths[column_index], cell.len)
		}
	}

	for row in table {
		for column_index, cell in row {
			if column_index != 0 {
				print(' ')
			}
			print(right_pad(cell, column_widths[column_index]))
		}
		println('')
	}
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

	edit_project(path)
}

fn edit(id string) {
	project_paths := list_project_paths()

	for index, project_path in project_paths {
		number := index + 1
		if number.str() == id {
			edit_project(project_path)
			return
		}
	}
}

fn edit_project(path string) {
	os.system([editor(), path].join(' '))
}

fn editor() string {
	return os.getenv_opt('EDITOR') or { panic(err) }
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

fn read_project(path string) Project {
	front_matter := retrieve_front_matter(path) or { panic(err) }

	mut project := Project{}

	for line in front_matter.split('\n') {
		key := line.all_before('=')
		value := line.all_after('=')

		if key == 'name' {
			project.name = value
		} else if key == 'date' {
			project.date = time.parse(value) or { panic(err) }
		}
	}

	return project
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

fn generate_slug(name string) string {
	mut slug := name

	slug = slug.to_lower()

	bad := 'àáâäæãåāăąçćčđďèéêëēėęěğǵḧîïíīįìıİłḿñńǹňôöòóœøōõőṕŕřßśšşșťțûüùúūǘůűųẃẍÿýžźż·/_,:;'
	good := 'aaaaaaaaaacccddeeeeeeeegghiiiiiiiilmnnnnoooooooooprrsssssttuuuuuuuuuwxyyzzz------'
	for i := 0; i < utf8.len(bad); i++ {
		bad_character := utf8.raw_index(bad, i)
		good_character := utf8.raw_index(good, i)
		for {
			found := slug.index(bad_character) or { -1 }
			if found > -1 {
				slug = slug[0..found] + good_character + slug[found + 1..]
			} else {
				break
			}
		}
	}

	mut white_space := regex.regex_opt(r'\s+') or { panic(err) }
	slug = white_space.replace_simple(slug, '-')

	slug = slug.replace('&', '-and-')

	mut non_words := regex.regex_opt(r'[^a-z0-9\-]') or { panic(err) }
	slug = non_words.replace_simple(slug, '')

	mut multiple_hyphens := regex.regex_opt(r'-{2,}') or { panic(err) }
	slug = multiple_hyphens.replace_simple(slug, '-')

	slug = slug.trim('-')

	return slug
}

fn list_project_paths() []string {
	file_type := 'md'

	return os.glob(os.join_path(os.home_dir(), '.projects', '*.' + file_type)) or { return [] }
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
		}
	}

	help()
}
