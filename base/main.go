package main

import "context"

func main() {
	as := apiService{
		Ctx: context.Background(),
	}

	as.depthWebsocket("bnbtc")
}