package main

import (
	"context"
)

type Ws_Binance interface {
	depthWebsocket(dsymbol string)
}

type apiService struct {
	Ctx context.Context
}
