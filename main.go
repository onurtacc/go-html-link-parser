package main

import (
	"flag"
	"fmt"
	"go-html-link-parser/parser"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	html := flag.String("html", "example.html", "the HTML doc to will be parsed")
	flag.Parse()

	t, err := ioutil.ReadFile(*html)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the HTML file : %s\n", *html))
	}
	r := strings.NewReader(string(t))
	links, err := parser.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}

func exit(msg string) {
	fmt.Printf(msg)
	os.Exit(1)
}
