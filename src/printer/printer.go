package printer

import (
	"fmt"
)

func UsageExample() {
	fmt.Println("USAGE: gogffu <url> <filepath>")
}

func DownloadStarted() {
	fmt.Println("Download started!")
}

func DownloadFinished() {
	fmt.Println("Download finished!")
}

func DownloadFromTo(url string, targetPath string) {
	fmt.Println("From:" + url)
	fmt.Println("To:" + targetPath)
}
