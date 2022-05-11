package markdown

import (
	md "github.com/golang-commonmark/markdown"
)

type Snippet struct {
	Content string
	Lang    string
}

func getSnippet(token md.Token) Snippet {
	switch token := token.(type) {
	case *md.CodeBlock:
		return Snippet{
			token.Content,
			"code",
		}
	case *md.CodeInline:
		return Snippet{
			token.Content,
			"code inline",
		}
	case *md.Fence:
		return Snippet{
			token.Content,
			token.Params,
		}
	}
	return Snippet{}
}

func FindSnippets(markdown string) []Snippet {
	snippets := []Snippet{}

	md := md.New(md.XHTMLOutput(true), md.Nofollow(true))
	tokens := md.Parse([]byte(markdown))

	for _, token := range tokens {
		snippets = append(snippets, getSnippet(token))
	}

	return snippets
}
