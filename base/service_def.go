package base

import (
	"context"
	"time"
	"github.com/go-kit/kit/log"
)

type Ws_Binance interface {
	depthWebsocket(dwr DepthWebsocketRequest) (chan *DepthEvent, chan struct{}, error)
}

type apiService struct {
	Logger log.Logger
	Ctx context.Context
}

type DepthWebsocketRequest struct {
	Symbol string
}

type OrderBook struct {
	LastUpdateID int `json:"lastUpdateId"`
}

type WSEvent struct {
	Type string
	Time time.Time
	Symbol string
}

type DepthEvent struct {
	WSEvent
	UpdateID int
}