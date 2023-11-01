package poller

import (
    "log"

    "web_listener/client"
    "web_listener/state"
)

type CompletedUrl struct {
    Url        string
    ErrorCount int // 为0表示正常访问到，否则就是访问不到
}

// Poller 完成如下工作：从in中读取url，并且通过http.Get发送请求获取数据；每当访问完成一个url，就会
// 将其封装成CompletedUrl发送到out中，便于后面继续重试访问。
// 而stateBuffer用于state信息的更新和打印（由StateMonitor完成）
func Poller(in <-chan *CompletedUrl, out chan<- *CompletedUrl, stateBuffer chan<- state.State) {
    // 阻塞读取in
    log.Printf("Poller start working.\n")

    for completeUrl := range in {

        resp, err := client.Conn.Get(completeUrl.Url)
        log.Printf("Poller read %q from chan and sent Http Request.\n", completeUrl.Url)
        var stateString string
        if err != nil {
            stateString = err.Error()
            completeUrl.ErrorCount++
        } else {
            stateString = resp.Status
            completeUrl.ErrorCount = 0
        }
        stateBuffer <- state.State{Url: completeUrl.Url, State: stateString}
        out <- completeUrl
    }
}
