package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"strings"
	"time"
)

var (
	webSocketKeepAlive = false
)

// 初始化连接
func intizlizeConnect() {
	webSocketKeepAlive = true
}

func (as *apiService) depthWebsocket(symbol string) {
	url := fmt.Sprintf("wss://stream.binance.com:9443/ws/%s@depth", strings.ToLower(symbol))

	c,_,err := websocket.DefaultDialer.Dial(url, nil)

	if err != nil {
		log.Fatal("dail err: ", err)
	}

	intizlizeConnect()

	go func() {
		defer c.Close()

		if webSocketKeepAlive {
			keepAlive(c, 15 * time.Second)
		}

		for {
			select {
			case <- as.Ctx.Done():
				return
			default:
				_, message, err := c.ReadMessage()

				if err != nil {
					log.Fatal("ws read err: ", err)
				}

				fmt.Println(message)
			}
		}
	}()
}


func keepAlive(c *websocket.Conn, timeout time.Duration) {
	ticker := time.NewTicker(timeout)

	lastResponse := time.Now()

	c.SetPongHandler(func(msg string) error {
		lastResponse = time.Now()
		return nil
	})

	go func() {
		defer ticker.Stop()

		for {
			deadline := time.Now().Add(10 * time.Second)
			err := c.WriteControl(websocket.PingMessage, []byte{}, deadline)
			if err != nil {
				return
			}

			<- ticker.C

			if time.Since(lastResponse) > timeout {
				c.Close()
				return
			}
		}
	}()
}