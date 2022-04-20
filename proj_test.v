module main

fn test_generate_slug() {
	assert generate_slug('SLUG') == 'slug'

	assert generate_slug('ãéǵ1') == 'aeg1'
	assert generate_slug('øø') == 'oo'

	assert generate_slug('slug    slug') == 'slug-slug'
	assert generate_slug('slug slug') == 'slug-slug'
	assert generate_slug('slug  slug') == 'slug-slug'

	assert generate_slug('and&and') == 'and-and-and'

	assert generate_slug('!!!') == ''

	assert generate_slug('multiple---hyphens') == 'multiple-hyphens'

	assert generate_slug('-start') == 'start'

	assert generate_slug('end-') == 'end'
}
