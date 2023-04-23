# gosublister

A fast subdomain enumerator for web URLs written in go with goroutines.

## Options
```
Usage:
  gosublister -u [URL] [Other Flags]

Flags:
  -u, --url string         The target domain [Required]
  -w, --wordlist string    Set the wordlist e.g shubs-domain.txt [Required]
  -d, --delay              Set delay in milliseconds per goroutine request
  -r, --response           Filter codes separated by space: e.g 200 302 402
  -t, --threads            Number of concurrent goroutine [Default=10]
  -h, --help               Display the help page

```
**Note:** If the target has rate limiting, consider using the time delay feature.

## Releases

A built version for Windows and Kali (Debian) has been created, please view **releases** page or check **release** folder.

## Demo Run
<img width="657" alt="ss" src="https://user-images.githubusercontent.com/28621928/233854467-43fc1dc5-b174-4110-9b94-9bccda58f7a1.png">

## Building Steps

Building gosublister only requires go-arg external module if you wish to re-build the project.

```
git clone https://github.com/n0mi1k/gosublister
go install https://github.com/alexflint/go-arg@latest
go build gosublister.go
```    