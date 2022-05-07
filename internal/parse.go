package internal

import (
	"log"
	"net/http"
	"net/url"

	"github.com/zhuima/mstoo/pkg"
)

func ParseUrl(urllink *url.URL, client *http.Client) (*pkg.Link, error) {
	req, err := http.NewRequest("GET", urllink.String(), nil)
	if err != nil {
		log.Printf("[ERROR] Can't Connect Url %s", err)
		return nil, err
	}

	// req.Close = true
	// Custom set header
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.3; WOW64; Trident/7.0; MAFSJS; rv:11.0) like Gecko")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[ERROR] Can't Get Response For Url %s", err)
		return nil, err

	}

	if resp != nil {
		defer resp.Body.Close()

	}

	// var link *pkg.Link

	// link = &pkg.Link{
	// 	Url:    urllink,
	// 	Status: resp.StatusCode,
	// }

	if resp.StatusCode >= pkg.HTTP_MIN_STATUS && resp.StatusCode <= pkg.HTTP_MAX_STATUS {
		link := pkg.NewLink(urllink, resp.StatusCode)
		return link, nil

	}

	return nil, err

}
