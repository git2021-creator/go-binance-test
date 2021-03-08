package main
import (
	"fmt"
	"github.com/adshao/go-binance/v2"
	"time"
)
func main() {
	wsDepthHandler := func(event *binance.WsDepthEvent) {
		fmt.Println(time.Now() ,"event -> : ",event)
	}
	errHandler := func(err error) {
		fmt.Println(time.Now() ,"err -> :", err)
	}
	doneC, stopC, err := binance.WsDepthServe("LTCBTC", wsDepthHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		time.Sleep(15 * time.Second)
		stopC <- struct{}{}
	}()
	<-doneC
}
