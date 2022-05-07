package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"

	"github.com/zhuima/mstoo/internal"
	"github.com/zhuima/mstoo/pkg"
)

func init() {
	log.SetOutput(os.Stdout)
}
func main() {

	urllistchan := make(chan *url.URL, 30)

	client := pkg.NewHttpRequest()

	go producer(urllistchan)

	start := time.Now()

	// 开启20个goroutine来处理
	noOfWorkers := 20
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go result(urllistchan, client, &wg)
	}
	wg.Wait()

	end := time.Now()
	log.Println("[INFO] End Parse Url", end.Sub(start))
}

func producer(urllistchan chan *url.URL) {
	// 同时允许10个并发
	file, err := internal.ReadFile("./urllist.txt")
	if err != nil {
		log.Printf("[ERROR] Can't Read File %s", err)
	}

	log.Println("[INFO] Start Parse Url")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// fmt.Println("[INFO] Parsing Url:", scanner.Text())
		// 每一个url检测都开启一个goroutine来检测
		reallink, err := url.Parse(scanner.Text())
		if err != nil {
			log.Printf("[ERROR] Can't Covert Url %s", err)
		}

		urllistchan <- reallink

	}
	close(urllistchan)
}

func result(urllistchan chan *url.URL, client *http.Client, wg *sync.WaitGroup) {
	// 循环channel读取结果
	count := 0
	var lck sync.Mutex

	lck.Lock()
	defer lck.Unlock()
	for v := range urllistchan {
		link, err := internal.ParseUrl(v, client)
		if err != nil {
			log.Printf("[ERROR] Can't Parse Url %s", err)
			// 出现异常就跳出当前循环
			continue
		}

		count++

		// log.Println("[INFO] Parse Url:", link)
		// 排除nil的情况
		if link != nil {
			fmt.Printf("|%-5d | %-150s| %-10d|\n", count, link.Url.String(), link.Status)

		}
	}
	wg.Done()
}
