package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/url"
	"os"
	"strings"
)

type Uri struct {
	query  string
	schema string
	host   string
	path   string
	new    bool
	params []string
}

const NEW = "new"
const OLD = "old"

func main() {

	var quietMode bool
	var encode bool
	var dryRun bool

	flag.BoolVar(&quietMode, "q", false, "quiet mode (no output)")
	flag.BoolVar(&encode, "e", false, "URL encode all query parameters")
	flag.BoolVar(&dryRun, "d", false, "don't append anything to file, just print new lines")
	flag.Parse()

	fn := flag.Arg(0)

	lines := make(map[string]Uri)

	var f io.WriteCloser

	if fn != "" {
		r, err := os.Open(fn)
		if err == nil {
			scanner := bufio.NewScanner(r)
			runScanner(*scanner, lines, quietMode, encode, OLD)
			r.Close()
		}
		if !dryRun {
			f, err = os.OpenFile(fn, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0644)
			if err != nil {
				fmt.Fprintf(os.Stderr, "failed to open file for writing: %s\n", err)
				return
			}
			defer f.Close()
		}
	}

	scanner := bufio.NewScanner(os.Stdin)
	runScanner(*scanner, lines, quietMode, encode, NEW)

	for _, element := range lines {
		if !quietMode {
			if element.new {
				if encode {
					fmt.Println(element.schema + "://" + element.host + element.path + "?" + url.QueryEscape(element.query))
				} else {
					fmt.Println(element.schema + "://" + element.host + element.path + "?" + strings.ReplaceAll(element.query, " ", "+"))
				}
			}
		}

		if !dryRun {
			if fn != "" {
				if encode {
					fmt.Fprintf(f, "%s\n", element.schema+"://"+element.host+element.path+"?"+url.QueryEscape(element.query))
				} else {
					fmt.Fprintf(f, "%s\n", element.schema+"://"+element.host+element.path+"?"+strings.ReplaceAll(element.query, " ", "+"))
				}
			}
		}
	}
}

func runScanner(scanner bufio.Scanner, lines map[string]Uri, quietMode bool, encode bool, new string) {

	var status bool
	if new == NEW {
		status = true
	} else {
		status = false
	}

	for scanner.Scan() {
		raw := scanner.Text()
		raw = strings.TrimSpace(raw)
		decodedUrl, err := url.QueryUnescape(raw)
		if err != nil {
			panic(err)
		}

		mUrl, err := url.Parse(decodedUrl)

		v, err := url.ParseQuery(mUrl.RawQuery)
		if err != nil {
			panic(err)
		}

		if len(v) > 0 {
			var params []string
			for k := range v {
				params = append(params, k)
			}

			hname := mUrl.Hostname
			name := fmt.Sprintln(hname() + mUrl.Path)
			p := Uri{mUrl.RawQuery, mUrl.Scheme, hname(), mUrl.Path, status, params}

			test := lines[name]
			if len(test.params) < len(v) {
				lines[name] = p
			} else {
				continue
			}
		}
	}
}
