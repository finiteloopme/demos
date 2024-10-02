package main

import (
	"context"
	"os"
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
	_url := os.Getenv("WS_URL")
	_apiKey := "test" //os.Getenv("API_KEY")
	// client, err := rpc.Dial("wss://" + _url + "?key=" + _apiKey)
	client, err := rpc.DialWebsocket(context.Background(), "wss://"+_url+"?key="+_apiKey, "*")
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Subscribe to new blocks.
	// clientSub, err := client.EthSubscribe(ctx, subch, "newHeads")
	// if err != nil {
	// 	log.Info("subscribe error:" + err.Error())
	// 	return
	// }

	// clientSub.

	// The connection is established now.
	// Update the channel with the current block.
	var lastBlock Block
	err := client.CallContext(ctx, &lastBlock, "eth_getBlockByNumber", "latest", false)
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
