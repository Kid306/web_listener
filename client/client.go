package client

import (
    "log"
    "net/http"
)

var Conn *Connector

type Connector struct {
    clients []*http.Client
}

func NewConnector(client ...*http.Client) *Connector {
    var clients []*http.Client
    if len(client) == 0 {
        log.Printf("No client is specified, using DefaultClient")
        clients = []*http.Client{http.DefaultClient}
    } else {
        clients = client
    }

    return &Connector{clients: clients}
}

func (c *Connector) Get(url string) (resp *http.Response, err error) {
    for i, client := range c.clients {
        resp, err = client.Get(url)
        if err == nil {
            // modify priority
            c.clients[0], c.clients[i] = c.clients[i], c.clients[0]
            return
        }
    }
    return
}
