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
    file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
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
        log.Fatal("[!] Error opening file")
    }
	return words
}


func enumSubdomain(url string, subdomain string, httpsFlag bool, codes []int, delay int) {
	suburl := ""
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

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
		if codes == nil {
			if resp.StatusCode >= 200 && resp.StatusCode < 500 && resp.StatusCode != 404 {
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


type args struct {
	URL string `arg:"-u, --url, required" help:"URL to enumerate [required]"`
	Wordlist string `arg:"-w, --wordlist, required" help:"Wordlist to use [required]"`
	Delay int `arg:"-d, --delay" default:"0" help:"Set delay in milliseconds for each goroutine request"`
	Response []int `arg:"-r" help:"Filter status codes separated by space: e.g 200 402 302"`
	Threads int `arg:"-t, --threads" default:"10" help:"Number of concurrent Goroutines"`
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

	if strings.Contains(url, "https://") {
		httpsFlag = true
	} else if strings.Contains(url, "http://") {
		httpsFlag = false
	} else {
		log.Fatal("[!] Please enter a valid URL specifying http:// or https://")
	}

	fmt.Println("[+] Starting GoSublister: " + url)
	fmt.Println("[+] Using Wordlist: " + wordfile)
	fmt.Println("[+] Thread Count: " + strconv.Itoa(threads))
	if delay != 0 {
		fmt.Println("[+] Using Delay: " + strconv.Itoa(delay) + "ms")
	}
	fmt.Println("[+] Found Subdomain URLs:")
	
	wordlist := readWordlist(wordfile)
	semaphore := make(chan struct{}, threads)

	for _, subdomain := range wordlist {
		semaphore <- struct{}{}

		go func(subdomain string) {
			enumSubdomain(url, subdomain, httpsFlag, codes, delay)
			<-semaphore
		}(subdomain)
	}

	for i := 0; i < cap(semaphore); i++ {
		semaphore <- struct{}{}
	}
}
