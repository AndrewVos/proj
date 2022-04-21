module main

import regex
import encoding.utf8
import math
import os

fn string_to_bool(s string) bool {
	return s == 'true'
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

fn editor() string {
	return os.getenv_opt('EDITOR') or { panic(err) }
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
