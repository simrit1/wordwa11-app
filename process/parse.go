package process

import (
	"strings"
    "encoding/json"
	"log"

    "github.com/pxtrez/wordwa11-app/utils"
)

func parse(data []byte) string {
    var thumbnail_url URL

    err := json.Unmarshal(data, &thumbnail_url)
    if err != nil {
        log.Fatal(err)
    }

    url := thumbnail_url.Thumbnail_url
    guid := strings.Split(url, "/")[len(strings.Split(url, "/")) - 1]
    parsed := "https://wordwall.net/create/editcontent?guid=" + guid

    return parsed
}