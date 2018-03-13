Package go-getfilefromurl
===================

Package go-getfilefromurl provides a very simple library for downloading file from given URL

## Installation
```shell
go get -u github.com/artdarek/go-getfilefromurl
```
or if your looking for the standalone client
```shell
go get -u github.com/artdarek/go-getfilefromurl/cmd/gogffu
```

## Examples

```go
package main

import (
	 "github.com/artdarek/go-getfilefromurl/src/downloader"
)

func main() {

    d := downloader.NewDownloader("http://www.domain.com/whatever/filename.zip", "/your/target/path/filename.zip")
    err := d.Start()
    if err == nil {
        fmt.Println("Download finished!")
    }    
    
}
```

## Contributing

Pull requests, bug fixes and issue reports are welcome.

Before proposing a change, please discuss your change by raising an issue.
