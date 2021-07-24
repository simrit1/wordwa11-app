package process

import (
	"io/ioutil"
    "log"
    "net/http"
)

func Fetch(link string) []byte {
    API_URL := "https://wordwall.net/api/oembed?url=" + link + "&format=json"
    
    resp, err := http.Get(API_URL)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
	
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
	
    return body
}