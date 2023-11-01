package client

import (
    "log"
    "net/http"
    "net/url"

    . "web_listener/constant"
)

func Prepare() {
    // ignore error
    uri, proxyError := url.Parse("http://localhost:7890")
    if proxyError != nil {
        log.Printf("---WARN: The proxy uri may be disabled: %q\n", uri)
    }

    var defaultClient, proxyClient *http.Client
    defaultClient = &http.Client{Timeout: ConnectionTimeout}
    proxyClient = &http.Client{Timeout: ConnectionTimeout, Transport: &http.Transport{Proxy: http.ProxyURL(uri)}}

    var clients []*http.Client

    // Make sure the client priority
    if proxyError != nil {
        clients = []*http.Client{defaultClient, proxyClient}
    } else {
        clients = []*http.Client{proxyClient, defaultClient}
    }

    Conn = NewConnector(clients...)
}
