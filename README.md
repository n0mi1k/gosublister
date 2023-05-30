# gosublister

A fast subdomain enumerator for web URLs written in go with goroutines.

## Options
```
Usage:
  gosublister -u [URL] [Other Flags]

Flags:
  -u, --url string         The target domain [Required]
  -w, --wordlist string    Set the wordlist e.g shubs-subdomains.txt [Required]
  -d, --delay              Set delay in ms per goroutine request [Default=100]
  -r, --response           Filter out codes separated by space: e.g 200 302 402
  -t, --threads            Number of concurrent goroutines [Default=10]
  -s, --timeout            Timeout for each request [Default=2s]
  -h, --help               Display the help page

```
**Note:** If the target has rate limiting, use the time delay feature and reduce threads.

## Releases

Built version for Kali (Debian), macOS (arm64, amd64) and Windows has been created, please view [releases](https://github.com/n0mi1k/gosublister/releases) page.

## Demo Run
<img width="700" alt="ss" src="https://github.com/n0mi1k/gosublister/assets/28621928/1b08992b-983a-4c4b-acf5-12615f1d91a4">

More subdomain wordlists can be found on seclists [here](https://github.com/danielmiessler/SecLists/tree/master/Discovery/DNS).

## Building Steps

Building gosublister only requires go-arg external module if you wish to re-build the project.

```
git clone https://github.com/n0mi1k/gosublister
go install https://github.com/alexflint/go-arg@latest
go build gosublister.go
```    