package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/zhuima/mstoo/internal"
	"github.com/zhuima/mstoo/pkg"
)

func init() {
	log.SetOutput(os.Stdout)
}
func main() {
	client := pkg.NewHttpRequest()

	file, err := internal.ReadFile("./urllist.txt")
	if err != nil {
		log.Printf("[ERROR] Can't Read File %s", err)
	}

	fmt.Println("[INFO] Start Parse Url")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	count := 0
	start := time.Now()
	for scanner.Scan() {
		// fmt.Println("[INFO] Parsing Url:", scanner.Text())
		reallink, err := url.Parse(scanner.Text())
		if err != nil {
			log.Printf("[ERROR] Can't Covert Url %s", err)
			continue
		}

		link, err := internal.ParseUrl(reallink, client)
		if err != nil {
			log.Printf("[ERROR] Can't Parse Url %s", err)
			continue
		}

		count++

		fmt.Printf("|%-5d | %-150s| %-10d|\n", count, link.Url.String(), link.Status)
	}

	end := time.Now()
	fmt.Println("cost time:", end.Sub(start))
}
