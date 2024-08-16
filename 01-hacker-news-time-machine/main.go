package main

/**
This uses goroutines to parse each individual athing
- The results don't keep the original order (for now)
- It fetches only the first page of hn (~30 titles)
**/

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

type Result struct {
	Index int
	Text  string
}

// Check if a Html node has a given class
func hasClass(node *html.Node, class string) bool {
	// For every attribute
	for _, a := range node.Attr {
		if a.Key == "class" && a.Val == class {
			return true
		}
	}
	return false
}

// Returns nth child of a node
func getNthChild(node *html.Node, index int) *html.Node {
	var i int
	// A for, starts at first node, goes to next sibling, until no more siblings
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		// If the sibling is a node, count it
		if c.Type == html.ElementNode {
			i++
			// If at the Nth child, return it
			if i == index {
				return c
			}
		}
	}

	return nil
}

func getTextOfNode(node *html.Node) string {
	// If it is a simple text node, return the data (it is the text here)
	if node.Type == html.TextNode {
		return node.Data
	}

	// Recursively traverse the child nodes and concatenate their texts
	var text string
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		text += getTextOfNode(c)
	}
	return text
}

func findAthingElements(n *html.Node, wg *sync.WaitGroup, results chan<- string) {
	if n.Type == html.ElementNode && hasClass(n, "athing") {
		// We found an element with the class "athing"

		// Increment the WaitGroup as we have another athing to process
		wg.Add(1)

		// Launch a gorouting to get the text inside the element
		// This is declared as an anonymous function
		go func(node *html.Node) {
			defer wg.Done() // Decrement the counter when the gorouting completes

			titleChild := getNthChild(node, 3)
			titleText := getTextOfNode(titleChild)
			if titleText != "" {
				// fmt.Println(titleText)
				results <- titleText // Send the result to the channel

			}
		}(n) // Actually call the goroutine
	}

	// Recursively traverse the HTML tree
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		findAthingElements(c, wg, results)
	}
}

func main() {

	// Fetch the page
	res, err := http.Get("https://news.ycombinator.com/")
	if err != nil {
		log.Fatal(err)
	}

	// Get the raw html
	content, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := html.Parse(strings.NewReader(string(content)))
	if err != nil {
		panic(err)
	}

	results := make(chan string)

	// Use a WaitGroup to wait for all goroutines to complete
	var wg sync.WaitGroup

	// var wg sync.WaitGroup
	go func() {
		wg.Wait()      // Wait for all goroutines to finish
		close(results) // Close the channel once all work is done
	}()

	// Traverse the HTML tree and find all elements with the class "athing"
	go findAthingElements(doc, &wg, results)

	// Wait for all goroutines to complete
	// wg.Wait()

	// Read from the results channel and print each result in order
	for result := range results {
		fmt.Println(result)
	}

}
