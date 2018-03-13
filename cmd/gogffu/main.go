package main

import (
    "os"
    "github.com/artdarek/go-getfilefromurl/src/printer"
    "github.com/artdarek/go-getfilefromurl/src/downloader"
)

func main() {

    if os.Args == nil || len(os.Args) <= 2 {
        printer.UsageExample()
        return
    }

    url := os.Args[1]
    file := os.Args[2]

    download(url, file)
}

func download(url string, targetPath string) {

    printer.DownloadStarted()
    printer.DownloadFromTo(url, targetPath)

    d := downloader.NewDownloader(url, targetPath)
    err := d.Start()
    if err == nil {
        printer.DownloadFinished()
    }
}

