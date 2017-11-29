package main

import (
    "fmt"
    "os"
    "github.com/artdarek/go-getfilefromurl/src/printer"
    "github.com/artdarek/go-getfilefromurl/src/downloader"
)

func main() {

    if (os.Args == nil || len(os.Args) <= 2) {
        printer.UsageExample()
        return
    }

    url := os.Args[1]
    file := os.Args[2]

    download(url, file)
}

func download(url string, targetPath string) {
    fmt.Println("Download started!")
    err := downloader.Start(targetPath, url)
    if (err == nil) {
        fmt.Println("Download finished!")
    }
}

