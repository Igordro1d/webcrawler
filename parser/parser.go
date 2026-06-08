package parser

import (
	"io"

	"golang.org/x/net/html"
)

// Extract reads HTML from r and returns every href found on <a> tags
func Extract(r io.Reader) ([]string, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	var links []string

	// walk visits a node, harvests it if it's a link, then recurses into its children
	var walk func(n *html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					links = append(links, attr.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}

	walk(doc)
	return links, nil
}
