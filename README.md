# Jeb
> Parse a list of URLs from stdin and output URLs that contain unique paths and query parameters. This is used to dedup URLs that have different values for query parameters. 

![Golang][go-lang]
![Bounty][bug-bounty]

Jeb is bug-bounty tool that takes stdin and outputs a list of URLs that can be fuzzed with other tools, such as Dalfox, Sqlmap, Commix, or Nuclei. If you frequently use Gau or Waybackurls, then there is a chance that the result of those tools will return duplicate URLs that vary slightly based on the value in the query parameters. This tool aims to elimiate the duplicate entries so that your tools aren't repetitively fuzzing the same URLs.

## Usage
```bash
> ./jeb -h 

Usage of ./jeb:
  -d    don't append anything to file, just print new lines
  -e    URL encode all query parameters
  -q    quiet mode (no output)
```

## Examples
```bash
# Take a list of urls and process with jeb
> cat urls.txt
https://example.com/someparam?x=1&y=2
https://example.com/someparam?x=1&y=2&z=4
https://example.com/someparam?x=abcd&y=7
https://example2.com/someparam?x=37&y=99
https://example2.com/someparam?x=efg&y=2&z=223

> cat urls.txt | jeb outfile.txt
https://example.com/someparam?x=1&y=2&z=4
https://example2.com/someparam?x=efg&y=2&z=223

# Do a waybackurl search and write unique URLs with query params to file
> cat domains.txt | waybackurls | jeb outfile.txt

# Do a gau search, pipe to jeb, and test the endopints with dalfox
> cat domains.txt | gau | jeb | dalfox pipe
```

If the output file jeb writes to contains existing values, those values will be read into jeb and only unique values will be displayed to stdout. 
```bash
> cat outfile.txt
https://example.com/someparam?x=1&y=2&z=4
https://example2.com/someparam?x=efg&y=2&z=223

> cat urls.txt
https://example3.com/someparam?x=efg&y=2&z=223

> cat urls.txt | jeb outfile.txt
https://example3.com/someparam?x=efg&y=2&z=223

> cat outfile.txt
https://example.com/someparam?x=1&y=2&z=4
https://example2.com/someparam?x=efg&y=2&z=223
https://example3.com/someparam?x=efg&y=2&z=223
```

## Flags
* To view output in stdout, but not append to the file, use the dry-run option ```-d```
* To append to the file, but not print anytyhing to stdout, use the quiet mode ```-q```

## Install
```
go install -v github.com/af001/jeb@latest
```

Or [download](https://github.com/af001/jeb/releases/) a binary release for your platform.

## References
Based on [anew](https://github.com/tomnomnom/anew) by [@tomnomnom](https://github.com/tomnomnom)

<!-- Markdown link & img dfn's -->
[go-lang]: https://img.shields.io/badge/Go-1.19-blue
[bug-bounty]: https://img.shields.io/badge/Bug-Bounty-red

