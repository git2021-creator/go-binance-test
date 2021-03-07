package base

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/websocket"
	"strings"
)

func (as *apiService) depthWebsocket(dwr DepthWebsocketRequest) (chan *DepthEvent, chan struct{}, error) {
	url := fmt.Sprintf("wss://stream.binance.com:9443/ws/%s@depth", strings.ToLower(dwr.Symbol))

	c,_,err := websocket.DefaultDialer.Dial(url, nil)

	if err != nil {
		log.Logger.Log("dail: ", err)
	}

	done := make(chan struct{})

	go func() {
		defer c.Close()
		defer close(done)

		for {
			select {
			case <- as.Ctx.Done():
				return
			default:
				_, message, err := c.ReadMessage()

				if err != nil {
					log.Logger.Log("ws read err: ", err)
				}

				fmt.Println(message)
			}
		}
	}()
	return nil, nil, nil
}

