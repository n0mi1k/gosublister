# gosublister

An uber fast subdomain enumerator for web URLs written in go using goroutines.

# Installation
To install, just run the below command or download pre-compiled binary from [release page](https://github.com/n0mi1k/gosublister/releases/tag/v1.1.0).
```
go install github.com/n0mi1k/gosublister@latest
```
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

## Demo Run
<img width="700" alt="ss" src="https://github.com/n0mi1k/gosublister/assets/28621928/1b08992b-983a-4c4b-acf5-12615f1d91a4">

More subdomain wordlists can be found on seclists [here](https://github.com/danielmiessler/SecLists/tree/master/Discovery/DNS).

