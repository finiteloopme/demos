package main

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
	log "github.com/finiteloopme/goutils/pkg/log"
)

type Block struct {
	Number *hexutil.Big
}

func main() {
	log.Info("In main")
	// Connect the client.
	_url := "ws.awm4nzzwx9ogrvhmbi1ezkz5j.blockchainnodeengine.com"
	_apiKey := "AIzaSyDbqUu36nne_ueHOqje7SmsFHbM_bItOaQ"
	client, err := rpc.Dial("wss://" + _url + "?key=" + _apiKey)
	if err != nil {
		log.Fatal(err)
	}

	subch := make(chan Block)
	go subscribeBlocks(client, subch)
	// Print events from the subscription as they arrive.
	for block := range subch {
		log.Info("latest block:" + block.Number.String())
	}

	log.Info("Conencting main")
}

// subscribeBlocks runs in its own goroutine and maintains
// a subscription for new blocks.
func subscribeBlocks(client *rpc.Client, subch chan Block) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Subscribe to new blocks.
	_, err := client.EthSubscribe(ctx, subch, "newHeads")
	if err != nil {
		log.Info("subscribe error:" + err.Error())
		return
	}

	// The connection is established now.
	// Update the channel with the current block.
	var lastBlock Block
	err = client.CallContext(ctx, &lastBlock, "eth_getBlockByNumber", "latest", false)
	if err != nil {
		log.Info("can't get latest block:" + err.Error())
		return
	}
	subch <- lastBlock

	// The subscription will deliver events to the channel. Wait for the
	// subscription to end for any reason, then loop around to re-establish
	// the connection.
	//log.Info("connection lost: " + string(<-sub.Err()))
}
