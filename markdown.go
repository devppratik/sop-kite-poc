package main

import (
	"bytes"
	"io"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func getGitHubMdURL(URL string) string {
	end := URL[len(URL)-3:]
	host := URL[8:18]
	if host == "github.com" && end == ".md" {
		newURL := "https://raw.githubusercontent.com" + URL[18:]
		newURL = strings.ReplaceAll(newURL, "/blob", "")
		return newURL
	}
	return URL
}

func convertMarkdownToHTML(body io.ReadCloser) io.Reader {
	md, err := io.ReadAll(body)
	if err != nil {
		panic(err)
	}
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	data := markdown.Render(doc, renderer)
	return bytes.NewReader(data)
}
