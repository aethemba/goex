package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var depth int

func main() {

	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
	}

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}

	forEachNode(doc, startElement, endElement)
	// outline(nil, doc)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) {
	var cont bool

	if pre != nil {
		cont = pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if cont == false {
			return
		} else {
			forEachNode(c, pre, post)
		}
	}

	if post != nil {
		cont = post(n)
	}
}

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {

		attrs := ""
		for _, att := range n.Attr {
			attrs += att.Key + "=" + "'" + att.Val + "' "
		}

		if n.FirstChild == nil {
			fmt.Printf("%*s<%s %s/>\n", depth*2, "", n.Data, attrs)
		} else {

			fmt.Printf("%*s<%s %s>\n", depth*2, "", n.Data, attrs)
		}
		depth++
	}

	if n.Type == html.TextNode {
		fmt.Printf("%*s%s\n", depth*2, "", n.Data)
	}

}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if n.FirstChild != nil {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}

}

func visit(links []string, n *html.Node) []string {

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	if n.Type == html.ElementNode && n.Data == "img" {
		for _, att := range n.Attr {
			if att.Key == "src" {
				links = append(links, att.Val)
			}
		}
	}

	if n.Type == html.ElementNode && n.Data == "script" {
		for _, att := range n.Attr {
			if att.Key == "src" {
				links = append(links, att.Val)
			}
		}
		// for _, att := range n.Attr {
		// 	if att.Key == "src" {
		// 		links = append(links, att.Val)
		// 	}
		// }
	}

	// fmt.Println(links)

	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}

	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}

	// for c := n.FirstChild; c != nil; c = c.NextSibling {
	// 	links = visit(links, c)
	// }
	return links
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}
