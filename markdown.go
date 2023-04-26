package main

import (
	"bytes"
	"io"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func getGitHubMdURL(URL string) (owner, repo, path string) {
	end := URL[len(URL)-3:]
	host := URL[8:18]
	if host == "github.com" && end == ".md" {
		owner, repo := getOwnerAndRepoName(URL)
		path := getReadmePath(URL)
		return owner, repo, path
	}
	return URL, "", ""
}

func getOwnerAndRepoName(str string) (owner, repo string) {
	startS, endS := ".com/", "/blob"
	s := strings.Index(str, startS)
	newS := str[s+len(startS):]
	e := strings.Index(newS, endS)
	result := newS[:e]
	combined := strings.SplitN(result, "/", -1)
	return combined[0], combined[1]
}

func getReadmePath(str string) string {
	e := strings.Index(str, "/blob/")
	result := str[e+len("/blob/"):]
	combined := strings.SplitN(result, "/", 2)
	return combined[1]
}

func convertMarkdownToHTML(body string) io.Reader {
	md := []byte(body)
	extensions := parser.CommonExtensions
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	data := markdown.Render(doc, renderer)
	return bytes.NewReader(data)
}
