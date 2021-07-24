package main

import (
    "fmt"

    "github.com/pxtrez/wordwa11-app/process"
    "github.com/toqueteos/webbrowser"
)

func main() {
    var link, foo string

    fmt.Print("[wordwa11] Test link: ")
    fmt.Scanln(&link)

    fmt.Println("[wordwa11] Fetching...")
    data := fetch(link)

    fmt.Println("[wordwa11] Parsing...")
    answer_page := parse(data)


    fmt.Println("[wordwa11] Answers can be found at: " + answer_page)
    fmt.Print("[wordwa11] Do you want to open answers in webbrowser? [y/n] ")
    fmt.Scanln(&foo)

    if strings.ToLower(foo) == "y" {
        fmt.Println("[wordwa11] Opening answers in webbrowser...")
        webbrowser.Open(answer_page)
    }
}