package main

import (
	"bufio"
	"embed"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"text/template"
	"time"
)

//go:embed template.html
var content embed.FS

type Item struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	Url      string `json:"url"`
	UnixTime int64  `json:"pubDate"`
	Content  string `json:"content"`
	PubDate  string
}

func (item *Item) ConvertUnixTime() {
	t := time.Unix(item.UnixTime, 0)
	item.PubDate = t.Format("2006-01-02 15:04:05")
}

type Payload struct {
	Items []Item
}

func main() {
	var inputLines []string
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputLines = append(inputLines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	var inputText string = strings.Join(inputLines, "")
	var items []Item

	if err := json.Unmarshal([]byte(inputText), &items); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	for i := range items {
		items[i].ConvertUnixTime()
	}

	payload := Payload{
		Items: items,
	}

	file, _ := content.ReadFile("template.html")
	tmpl, _ := template.New("").Parse(string(file))
	tmpl.Execute(os.Stdout, payload)
}
