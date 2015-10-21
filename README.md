# go-hn

`go-hn` is a hacker news client for the terminal.  There's not much to this program code-wise, but this was a another happy surprise from GO as the XML parsing worked the first time.

Quick install: `go get github.com/sspencer/go-hn`

## Features

* Grabs latest hacker news headlines from the Hacker News [RSS Feed](https://news.ycombinator.com/rss) and displays them in the terminal
* Opens desired article in default browser

![Hacker News Terminal](screenshot.png?raw=true "Hacker News Terminal Screenshot")

## Dependencies

* Working Go environment
* utf-8 terminal with colors

## Setup

* Install the `go-hn` binary into your $GOPATH with `go get github.com/sspencer/go-hn`
* Invoked `go-hn` (assuming go/bin is in your path)
* If you prefer an executable named `hn` or just want to play and reinstall the code, make sure `$GOBIN` is set, `cd $GOPATH/src/github.com/sspencer/go-hn` and type `go install hn.go`.

## Parsing XML

GO makes it SO easy!  Just define a few types and away you GO.

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

    body, err := ioutil.ReadAll(resp.Body)
    ...
    var v Result
    err = xml.Unmarshal(body, &v)

