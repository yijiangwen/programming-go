// Usage:   findlinks3 [url]
// Example: findlinks3 http://golang.org
package main

import (
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

func main() {
	// Crawl the web breadth first.
	// Starting from the command line arguments.
	breadthFirst(crawl, os.Args[1:])
}

func crawl(url string) []string {
	fmt.Println(url)

	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}

	return list
}

// breadthFirst calls f for each item in the worklist/
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}
