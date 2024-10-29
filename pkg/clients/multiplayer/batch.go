// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package multiplayer

import (
	"context"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/clients/multiplayer/jsonrpc"
)

type RequestRPC struct {
	retHandler rpcCallback
	rpcRequest *jsonrpc.RequestRPC
}

type rpcCallback func(err error, response *jsonrpc.ResponseRPC)

func (cli *ClientJsonRPC) Batch(ctx context.Context, requests ...RequestRPC) {

	var rpcRequests jsonrpc.RequestsRPC
	callbacks := make(map[jsonrpc.ID]rpcCallback)
	for _, request := range requests {
		rpcRequests = append(rpcRequests, request.rpcRequest)
		callbacks[request.rpcRequest.ID] = request.retHandler
	}
	var err error
	var rpcResponses jsonrpc.ResponsesRPC
	rpcResponses, err = cli.rpc.CallBatch(ctx, rpcRequests)
	for id, response := range rpcResponses.AsMap() {
		if callback := callbacks[id]; callback != nil {
			callback(err, response)
		}
	}
}
