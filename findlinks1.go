package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var m = make(map[string]int)

func main() {

	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
	}

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}

	words, images := countWordsAndImages(doc)

	fmt.Println("---")
	fmt.Println("Words: ", words)
	fmt.Println("Images: ", images)
	// var m = make(map[string]int)
	// m = tagFreq(m, doc)

	// fmt.Println("\n\nMap contents:")

	// for k, v := range m {
	// 	fmt.Printf("%s: %d\n", k, v)
	// }

	//fmt.Println("\n\nPrinting outline")
	//outline(nil, doc)
	//textNode(doc)
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Data == "script" || n.Data == "style" {
		return
	}

	if n.Type == html.TextNode {
		words = len(strings.Split(n.Data, " "))
	}

	if n.Type == html.ElementNode && n.Data == "img" {
		images = 1
	}

	var w, i int
	if n.FirstChild != nil {
		w, i = countWordsAndImages(n.FirstChild)
		// fmt.Println(w, i)
		words += w
		images += i
	}

	if n.NextSibling != nil {
		w, i = countWordsAndImages(n.NextSibling)
		// fmt.Println(w, i)
		words += w
		images += i
	}

	return words, images
}

func tagFreq(k map[string]int, n *html.Node) map[string]int {

	if n.Type == html.ElementNode {
		k[n.Data]++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		k = tagFreq(k, c)
	}

	return k
}

func textNode(n *html.Node) {

	if n.Data == "script" || n.Data == "style" {
		fmt.Println("SCRIPT OR STYLE NODE!", n.Data, n.Type)
		t := html.TextNode == n.Type
		fmt.Println("Is text node: ", t)
		return
	}

	if n.Type == html.TextNode && n.Data != "script" && n.Data != "style" {
		fmt.Printf("Data: %s\n", n.Data)
	}

	if n.FirstChild != nil {
		textNode(n.FirstChild)
	}

	if n.NextSibling != nil {
		textNode(n.NextSibling)
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
