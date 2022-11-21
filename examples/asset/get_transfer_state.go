package main

import (
	"log"

	"github.com/iaping/go-okx/examples"
	"github.com/iaping/go-okx/rest/api/asset"
)

func main() {
	param := &asset.GetTransferStateParam{
		TransId: "123456",
	}
	req, resp := asset.NewGetTransferState(param)
	if err := examples.Client.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*asset.GetTransferStateResponse))
}
