package main

import (
    "log"
    "time"

    "web_listener/client"
    "web_listener/poller"

    "web_listener/state"

    . "web_listener/constant"
)

var urls []string

func Recheck(completedUrl *poller.CompletedUrl, pendingUrls chan<- *poller.CompletedUrl) {
    time.Sleep(RecheckBaseInterval + time.Duration(completedUrl.ErrorCount)*RecheckTimeoutUnit)
    log.Printf("Recheck url: %q\n", completedUrl.Url)
    pendingUrls <- completedUrl
}

func PrepareUrls() {
    urls = []string{
        "https://www.baidu.com",
        "https://www.google.com/",
        "https://golang.org/",
        "https://blog.golang.org/",
    }
}

func main() {
    Prepare()

    pendingUrls := make(chan *poller.CompletedUrl)
    completedUrls := make(chan *poller.CompletedUrl)

    // 监听stateOuter中的数据，并且周期性的打印与state有关的信息
    stateBuffer, _ := state.Monitor(LookInterval)
    // 单独开一个协程来读取pending中的url，并且http结果发送到stateOuter中
    go poller.Poller(pendingUrls, completedUrls, stateBuffer)

    // 初始化pending
    go func() {
        for _, u := range urls {
            log.Printf("Initial url: %q\n", u)
            pendingUrls <- &poller.CompletedUrl{Url: u}
        }
    }()

    // recheck: 主要作用就是将已经访问过的url再周期性的访问，来实现持续更新
    // 否则就只会根据“初始化pending”中的url访问一次
    // 注意下方的代码不会退出
    for completedUrl := range completedUrls {
        go Recheck(completedUrl, pendingUrls)
    }
}

func Prepare() {
    PrepareUrls()

    client.Prepare()
}
