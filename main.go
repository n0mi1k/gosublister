package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
	"net/http"
	"strconv"
	"crypto/tls"
	"github.com/alexflint/go-arg"
	"time"
)


func readWordlist(filename string) []string {
	logger := log.New(os.Stderr, "", 0)
    file, err := os.Open(filename)
    if err != nil {
        logger.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
	const maxCapacity int = 20000000
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	words := make([]string, 0)
    for scanner.Scan() {
		words = append(words, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        logger.Fatal("[!] Error opening file")
    }
	return words
}


func enumSubdomain(url string, subdomain string, httpsFlag bool, codes []int, delay int, timeout int) {
	suburl := ""
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
		Timeout: time.Duration(timeout) * time.Second,
	}

	if delay > 0 {
		time.Sleep(time.Duration(delay) * time.Millisecond)
	} 
	
	if httpsFlag {
		if strings.Contains(url, "www.") {
			suburl = strings.Replace(url, "https://www.", "https://" + subdomain + ".", 1)
		} else {
			suburl = strings.Replace(url, "https://", "https://" + subdomain + ".", 1)
		}
	} else {
		if strings.Contains(url, "www.") {
			suburl = strings.Replace(url, "http://www.", "http://" + subdomain + ".", 1)
		} else {
			suburl = strings.Replace(url, "http://", "http://" + subdomain + ".", 1)
		}
	}
	
	resp, err := client.Get(suburl)
	if err == nil {
		if resp.StatusCode == 429 {
			for resp.StatusCode == 429 {
				fmt.Println("[!!!!] Rate Limited.. Reduce Threads, Increase Delay! Sleep 30s... ")
				time.Sleep(30 * time.Second)
				resp, err = client.Get(suburl)
			}
		} else {
			if codes == nil {
				if resp.StatusCode >= 200 && resp.StatusCode <= 599 && resp.StatusCode != 429 {
					fmt.Println(suburl + " [" + strconv.Itoa(resp.StatusCode) + "]")
				}
			} else {
				for _, code := range codes {
					if resp.StatusCode == code {
						fmt.Println(suburl + " [" + strconv.Itoa(resp.StatusCode) + "]")
					}
				}
			}
		}
	}
}


type args struct {
	URL string `arg:"-u, --url, required" help:"URL to enumerate [required]"`
	Wordlist string `arg:"-w, --wordlist, required" help:"Wordlist to use [required]"`
	Delay int `arg:"-d, --delay" default:"100" help:"Set delay in milliseconds for each goroutine request"`
	Response []int `arg:"-r" help:"Filter status codes separated by space: e.g 200 402 302"`
	Threads int `arg:"-t, --threads" default:"10" help:"Number of concurrent Goroutines"`
	Timeout int `arg:"-s, --timeout" default:"2" help:"Set timeout in seconds for each request"`
}


func (args) Description() string {
	return "This program enumerates subdomain URLs quickly using the power of Goroutines."
}


func main() {
	var httpsFlag bool = false
	var args args 
	arg.MustParse(&args)

	url := args.URL
	codes := args.Response
	threads := args.Threads	
	wordfile :=	args.Wordlist
	delay := args.Delay	
	timeout := args.Timeout

	logger := log.New(os.Stderr, "", 0)

	if strings.Contains(url, "https://") {
		httpsFlag = true
	} else if strings.Contains(url, "http://") {
		httpsFlag = false
	} else {
		logger.Fatal("[!] Enter a valid URL specifying http:// or https://")
	}
	
	wordlist := readWordlist(wordfile)

	fmt.Println("[*] Starting GoSublister: " + url)
	fmt.Println("[+] Using Wordlist: " + wordfile)
	fmt.Println("[+] Request Timeout: " + strconv.Itoa(timeout) + "s")
	fmt.Println("[+] Thread Count: " + strconv.Itoa(threads))
	fmt.Println("[+] Using Delay: " + strconv.Itoa(delay) + "ms")

	if codes != nil {
		message := fmt.Sprintf("[+] Ignored Codes: %s", strings.Join(strings.Fields(fmt.Sprint(codes)), ", "))
		fmt.Println(message)
	}

	fmt.Println("[+] Found Subdomain URLs:")
	
	
	semaphore := make(chan struct{}, threads)

	for _, subdomain := range wordlist {
		semaphore <- struct{}{}

		go func(subdomain string) {
			enumSubdomain(url, subdomain, httpsFlag, codes, delay, timeout)
			<-semaphore
		}(subdomain)
	}

	for i := 0; i < cap(semaphore); i++ {
		semaphore <- struct{}{}
	}
}
