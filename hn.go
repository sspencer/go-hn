package main

import (
	"encoding/xml"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/skratchdot/open-golang/open"
)

type Item struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
}

type Channel struct {
	Items []Item `xml:"item"`
}

type Result struct {
	Channel Channel `xml:"channel"`
}

func fetchRSS(url string) []Item {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var v Result
	err = xml.Unmarshal(body, &v)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return v.Channel.Items
}

func displayRSS(items []Item) {
	for i, item := range items {
		index := i + 1
		color.Set(color.FgRed)
		fmt.Printf("%2d.", index)
		color.Unset()
		fmt.Printf(" %s\n", html.UnescapeString(item.Title))
	}
}

func bound(i, lower, upper int) int {
	if i < lower {
		return lower
	} else if i > upper {
		return upper
	} else {
		return i
	}
}

func main() {
	var items []Item
	var input string
	refresh := true
	var prompt = color.New(color.FgBlue).PrintfFunc()

	for {
		if refresh {
			items = fetchRSS("https://news.ycombinator.com/rss")
			displayRSS(items)
			refresh = false
		}

		prompt("Type post number to open, (r) to refresh, (q) to quit: ")
		fmt.Scanln(&input)

		if strings.ToLower(input) == "r" {
			refresh = true
			continue
		}

		if strings.ToLower(input) == "q" {
			fmt.Println("Good bye!")
			os.Exit(0)
		}

		i, err := strconv.Atoi(input)
		if err != nil {
			color.Yellow("Try again")
			continue
		}

		i = bound(i, 1, len(items))
		open.Run(items[i-1].Link) // open in default browser

		fmt.Println()
		displayRSS(items)
	}
}
