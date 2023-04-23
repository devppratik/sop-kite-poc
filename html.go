package main

import (
	"fmt"
	"net/http"

	"github.com/rivo/tview"
	"golang.org/x/net/html"
)

var numLinks int = 0

func traverseHTMLDoc(n *html.Node, textView *tview.TextView) int {
	if n.Type == html.ElementNode {
		switch n.Data {
		case "a":
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					fmt.Fprintf(textView, `["%d"]%s[""]`, numLinks, attr.Val)
					numLinks++
				}
			}
		case "img":
			for _, attr := range n.Attr {
				if attr.Key == "src" {
					fmt.Fprintf(textView, `["%d"]%s[""]`, numLinks, attr.Val)
					numLinks++
				}
			}
		}
	} else if n.Type == html.TextNode {
		fmt.Fprintf(textView, "%s", n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		traverseHTMLDoc(c, textView)
	}
	return numLinks
}

func fetchHTMLContent(URL string, textView *tview.TextView) {
	textView.Clear()
	numLinks = 0
	rawURL := getGitHubMdURL(URL)
	resp, err := http.Get(rawURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body := convertMarkdownToHTML(resp.Body)

	// Parse the HTML file
	doc, err := html.Parse(body)
	if err != nil {
		panic(err)
	}
	traverseHTMLDoc(doc, textView)
}
